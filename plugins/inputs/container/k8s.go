package container

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
)

const (
	defaultServiceAccountPath = "/run/secrets/kubernetes.io/serviceaccount/token"
)

func buildPodMetrics(summaryApi *SummaryMetrics, dropTags []string, ignorePodNameRegexps []*regexp.Regexp, category string) ([]*io.Point, error) {
	var pts []*io.Point

	for _, pod := range summaryApi.Pods {
		if len(pod.Containers) == 0 {
			continue
		}
		tags := map[string]string{
			"node_name": summaryApi.Node.NodeName,
			"namespace": pod.PodRef.Namespace,
		}
		if !RegexpMatchString(ignorePodNameRegexps, pod.PodRef.Name) {
			tags["pod_name"] = pod.PodRef.Name
		}
		if category == "object" {
			tags["name"] = pod.PodRef.UID
		}
		for _, key := range dropTags {
			if _, ok := tags[key]; ok {
				delete(tags, key)
			}
		}

		fields := make(map[string]interface{})
		fields["cpu_usage_nanocores"] = float64(pod.CPU.UsageNanoCores)
		fields["cpu_usage_core_nanoseconds"] = float64(pod.CPU.UsageCoreNanoSeconds)
		fields["memory_usage_bytes"] = float64(pod.Memory.UsageBytes)
		fields["memory_working_set_bytes"] = float64(pod.Memory.WorkingSetBytes)
		fields["memory_rss_bytes"] = float64(pod.Memory.RSSBytes)
		fields["memory_page_faults"] = float64(pod.Memory.PageFaults)
		fields["memory_major_page_faults"] = float64(pod.Memory.MajorPageFaults)
		fields["network_rx_bytes"] = float64(pod.Network.RXBytes())
		fields["network_rx_errors"] = float64(pod.Network.RXErrors())
		fields["network_tx_bytes"] = float64(pod.Network.TXBytes())
		fields["network_tx_errors"] = float64(pod.Network.TXErrors())

		if cpuPrecent, err := pod.CPU.Percent(); err == nil {
			fields["cpu_usage"] = cpuPrecent
		}

		if category == "object" {
			message, err := json.Marshal(func() map[string]interface{} {
				var result = make(map[string]interface{}, len(tags)+len(fields))
				for k, v := range tags {
					result[k] = v
				}
				for k, v := range fields {
					result[k] = v
				}
				return result
			}())
			if err != nil {
				l.Warnf("json marshal failed: %s", err)
			} else {
				fields["message"] = string(message)
			}
		}

		pt, err := io.MakePoint(kubeletPodName, tags, fields, time.Now())
		if err != nil {
			return nil, err
		}
		pts = append(pts, pt)
	}

	return pts, nil
}

// Kubernetes represents the config object for the plugin
type Kubernetes struct {
	URL string `toml:"kubelet_url"`
	// Bearer Token authorization file path
	BearerToken       string   `toml:"bearer_token"`
	BearerTokenString string   `toml:"bearer_token_string"`
	IgnorePodName     []string `toml:"ignore_pod_name"`
	ClientConfig

	ignorePodNameRegexps []*regexp.Regexp
	roundTripper         http.RoundTripper
}

func (k *Kubernetes) Init() error {
	// If neither are provided, use the default service account.
	if k.BearerToken == "" && k.BearerTokenString == "" {
		k.BearerToken = defaultServiceAccountPath
	}

	if k.BearerToken != "" {
		token, err := ioutil.ReadFile(k.BearerToken)
		if err != nil {
			return err
		}
		k.BearerTokenString = strings.TrimSpace(string(token))
	}

	// reset
	k.ignorePodNameRegexps = k.ignorePodNameRegexps[:0]

	for _, n := range k.IgnorePodName {
		re, err := regexp.Compile(n)
		if err != nil {
			return err
		}
		k.ignorePodNameRegexps = append(k.ignorePodNameRegexps, re)
	}

	return nil
}

func (k *Kubernetes) GatherPodMetrics(dropTags []string, category string) ([]*io.Point, error) {
	summaryApi, err := k.GetSummaryMetrics()
	if err != nil {
		return nil, err
	}
	return buildPodMetrics(summaryApi, dropTags, k.ignorePodNameRegexps, category)
}

func (k *Kubernetes) GatherPodInfo(containerID string) (map[string]string, error) {
	podApi, err := k.GetPods()
	if err != nil {
		return nil, err
	}

	containerID = fmt.Sprintf("docker://%s", containerID)
	var m = make(map[string]string)

	for _, podMetadata := range podApi.Items {
		if len(podMetadata.Status.ContainerStatuses) == 0 {
			continue
		}
		for _, containerStauts := range podMetadata.Status.ContainerStatuses {
			if containerStauts.ContainerID == containerID {
				m["pod_name"] = podMetadata.Metadata.Name
				m["pod_namespace"] = podMetadata.Metadata.Namespace
				break
			}
		}
	}

	return m, nil
}

func (k *Kubernetes) GatherPodUID(containerID string) (string, error) {
	podApi, err := k.GetPods()
	if err != nil {
		return "", err
	}

	containerID = fmt.Sprintf("docker://%s", containerID)

	for _, podMetadata := range podApi.Items {
		if len(podMetadata.Status.ContainerStatuses) == 0 {
			continue
		}
		for _, containerStauts := range podMetadata.Status.ContainerStatuses {
			if containerStauts.ContainerID == containerID {
				return podMetadata.Metadata.UID, nil
			}
		}
	}

	return "", nil
}

func (k *Kubernetes) GatherWorkName(uid string) (string, error) {
	statsSummaryApi, err := k.GetSummaryMetrics()
	if err != nil {
		return "", err
	}

	for _, podMetadata := range statsSummaryApi.Pods {
		if len(podMetadata.Containers) == 0 {
			continue
		}

		if podMetadata.PodRef.UID == uid {
			return podMetadata.Containers[0].Name, nil
		}
	}

	return "", nil
}

func (k *Kubernetes) GatherNodeName(stats *SummaryMetrics) string {
	return stats.Node.NodeName
}

func (k *Kubernetes) GetPods() (*Pods, error) {
	var podApi Pods
	err := k.LoadJson(fmt.Sprintf("%s/pods", k.URL), &podApi)
	if err != nil {
		return nil, err
	}
	return &podApi, nil
}

func (k *Kubernetes) GetSummaryMetrics() (*SummaryMetrics, error) {
	var summaryApi SummaryMetrics
	err := k.LoadJson(fmt.Sprintf("%s/stats/summary", k.URL), &summaryApi)
	if err != nil {
		return nil, err
	}
	return &summaryApi, err
}

func (k *Kubernetes) LoadJson(url string, v interface{}) error {
	var req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	var resp *http.Response
	tlsCfg, err := k.ClientConfig.TLSConfig()
	if err != nil {
		return err
	}
	if k.roundTripper == nil {
		k.roundTripper = &http.Transport{
			TLSHandshakeTimeout:   5 * time.Second,
			TLSClientConfig:       tlsCfg,
			ResponseHeaderTimeout: 5 * time.Second,
		}
	}
	req.Header.Set("Authorization", "Bearer "+k.BearerTokenString)
	req.Header.Add("Accept", "application/json")

	resp, err = k.roundTripper.RoundTrip(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request to %s: %s", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s returned HTTP status %s", url, resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return fmt.Errorf(`Error parsing response: %s`, err)
	}

	return nil
}

type Pods struct {
	Kind       string    `json:"kind"`
	ApiVersion string    `json:"apiVersion"`
	Items      []PodItem `json:"items"`
}

type PodItem struct {
	Metadata PodItemMetadata `json:"metadata"`
	Status   PodItemStatus   `json:"status"`
}

type PodItemMetadata struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	UID       string            `json:"uid"`
	Labels    map[string]string `json:"labels"`
}

type PodItemStatus struct {
	ContainerStatuses []struct {
		ContainerID  string `json:"containerID"`
		RestartCount int64  `json:"restartCount"`
	} `json:"containerStatuses"`
}

type SummaryMetrics struct {
	Node NodeMetrics  `json:"node"`
	Pods []PodMetrics `json:"pods"`
}

type NodeMetrics struct {
	NodeName         string             `json:"nodeName"`
	SystemContainers []ContainerMetrics `json:"systemContainers"`
	CPU              CPUMetrics         `json:"cpu"`
	Memory           MemoryMetrics      `json:"memory"`
	Network          NetworkMetrics     `json:"network"`
	FileSystem       FileSystemMetrics  `json:"fs"`
	Runtime          RuntimeMetrics     `json:"runtime"`
}

type ContainerMetrics struct {
	Name   string            `json:"name"`
	CPU    CPUMetrics        `json:"cpu"`
	Memory MemoryMetrics     `json:"memory"`
	RootFS FileSystemMetrics `json:"rootfs"`
	LogsFS FileSystemMetrics `json:"logs"`
}

type RuntimeMetrics struct {
	ImageFileSystem FileSystemMetrics `json:"imageFs"`
}

type CPUMetrics struct {
	Time                 time.Time `json:"time"`
	UsageNanoCores       int64     `json:"usageNanoCores"`
	UsageCoreNanoSeconds int64     `json:"usageCoreNanoSeconds"`
}

func (c *CPUMetrics) Percent() (float64, error) {
	if c.UsageNanoCores == 0 {
		return -1, fmt.Errorf("cpu usageNanoCores cannot be zero")

	}
	// source link: https://github.com/kubernetes/heapster/issues/650#issuecomment-147795824
	// cpu_usage_core_nanoseconds / (cpu_usage_nanocores * 1000000000) * 100
	return float64(c.UsageCoreNanoSeconds) / float64(c.UsageNanoCores*1000000000) * 100, nil
}

type PodMetrics struct {
	PodRef     PodReference       `json:"podRef"`
	StartTime  *time.Time         `json:"startTime"`
	Containers []ContainerMetrics `json:"containers"`
	CPU        CPUMetrics         `json:"cpu"`
	Memory     MemoryMetrics      `json:"memory"`
	Network    NetworkMetrics     `json:"network"`
}

type PodReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	UID       string `json:"uid"`
}

type MemoryMetrics struct {
	Time            time.Time `json:"time"`
	AvailableBytes  int64     `json:"availableBytes"`
	UsageBytes      int64     `json:"usageBytes"`
	WorkingSetBytes int64     `json:"workingSetBytes"`
	RSSBytes        int64     `json:"rssBytes"`
	PageFaults      int64     `json:"pageFaults"`
	MajorPageFaults int64     `json:"majorPageFaults"`
}

func (m *MemoryMetrics) Percent() (float64, error) {
	if m.AvailableBytes+m.UsageBytes == 0 {
		return -1, fmt.Errorf("memory total cannot be zero")
	}
	// mem_usage_percent = memory_usage_bytes / (memory_usage_bytes + memory_available_bytes)
	return float64(m.UsageBytes) / float64(m.UsageBytes+m.AvailableBytes), nil
}

type FileSystemMetrics struct {
	AvailableBytes int64 `json:"availableBytes"`
	CapacityBytes  int64 `json:"capacityBytes"`
	UsedBytes      int64 `json:"usedBytes"`
}

type NetworkMetrics struct {
	Time       time.Time `json:"time"`
	Interfaces []struct {
		Name     string `json:"name"`
		RXBytes  int64  `json:"rxBytes"`
		RXErrors int64  `json:"rxErrors"`
		TXBytes  int64  `json:"txBytes"`
		TXErrors int64  `json:"txErrors"`
	} `json:"interfaces"`
}

func (n NetworkMetrics) RXBytes() int64 {
	var sum int64
	for _, i := range n.Interfaces {
		sum += i.RXBytes
	}
	return sum
}

func (n NetworkMetrics) RXErrors() int64 {
	var sum int64
	for _, i := range n.Interfaces {
		sum += i.RXErrors
	}
	return sum
}

func (n NetworkMetrics) TXBytes() int64 {
	var sum int64
	for _, i := range n.Interfaces {
		sum += i.TXBytes
	}
	return sum
}

func (n NetworkMetrics) TXErrors() int64 {
	var sum int64
	for _, i := range n.Interfaces {
		sum += i.TXErrors
	}
	return sum
}

type VolumeMetrics struct {
	Name           string `json:"name"`
	AvailableBytes int64  `json:"availableBytes"`
	CapacityBytes  int64  `json:"capacityBytes"`
	UsedBytes      int64  `json:"usedBytes"`
}
