// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package container collect container metrics/loggings/objects.
package container

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/GuanceCloud/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/plugins/inputs"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/tailer"

	dkio "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io"
	dkpt "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io/point"
)

var (
	_ inputs.ReadEnv   = (*Input)(nil)
	_ inputs.Singleton = (*Input)(nil)
)

type Input struct {
	Endpoints                   []string `toml:"endpoints"`
	DeprecatedDockerEndpoint    string   `toml:"docker_endpoint"`
	DeprecatedContainerdAddress string   `toml:"containerd_address"`

	EnableContainerMetric       bool `toml:"enable_container_metric"`
	EnableK8sMetric             bool `toml:"enable_k8s_metric"`
	EnablePodMetric             bool `toml:"enable_pod_metric"`
	EnableK8sEvent              bool `toml:"enable_k8s_event"`
	Election                    bool `toml:"election"`
	EnableExtractK8sLabelAsTags bool `toml:"extract_k8s_label_as_tags"`

	K8sURL                                            string `toml:"kubernetes_url"`
	K8sBearerToken                                    string `toml:"bearer_token"`
	K8sBearerTokenString                              string `toml:"bearer_token_string"`
	EnableAutoDiscoveryOfPrometheusPodAnnotations     bool   `toml:"enable_auto_discovery_of_prometheus_pod_annotations"`
	EnableAutoDiscoveryOfPrometheusServiceAnnotations bool   `toml:"enable_auto_discovery_of_prometheus_service_annotations"`
	EnableAutoDiscoveryOfPrometheusPodMonitors        bool   `toml:"enable_auto_discovery_of_prometheus_pod_monitors"`
	EnableAutoDiscoveryOfPrometheusServiceMonitors    bool   `toml:"enable_auto_discovery_of_prometheus_service_monitors"`

	ContainerIncludeLog               []string          `toml:"container_include_log"`
	ContainerExcludeLog               []string          `toml:"container_exclude_log"`
	LoggingExtraSourceMap             map[string]string `toml:"logging_extra_source_map"`
	LoggingSourceMultilineMap         map[string]string `toml:"logging_source_multiline_map"`
	LoggingAutoMultilineDetection     bool              `toml:"logging_auto_multiline_detection"`
	LoggingAutoMultilineExtraPatterns []string          `toml:"logging_auto_multiline_extra_patterns"`
	LoggingSearchInterval             time.Duration     `toml:"logging_search_interval"`
	LoggingMinFlushInterval           time.Duration     `toml:"logging_min_flush_nterval"`
	LoggingMaxMultilineLifeDuration   time.Duration     `toml:"logging_max_multiline_life_duration"`
	LoggingRemoveAnsiEscapeCodes      bool              `toml:"logging_remove_ansi_escape_codes"`

	Tags map[string]string `toml:"tags"`
	DeprecatedConf

	Feeder dkio.Feeder
	Tagger dkpt.GlobalTagger

	semStop *cliutils.Sem // start stop signal
	pause   *atomic.Bool
	chPause chan bool
}

func (*Input) SampleConfig() string { return sampleCfg }

func (*Input) Catalog() string { return "container" }

func (*Input) PipelineConfig() map[string]string { return nil }

func (*Input) GetPipeline() []*tailer.Option { return nil }

func (*Input) RunPipeline() { /*nil*/ }

func (*Input) Singleton() { /*nil*/ }

func (*Input) AvailableArchs() []string {
	return []string{datakit.OSLabelLinux, datakit.LabelK8s, datakit.LabelDocker}
}

func (*Input) SampleMeasurement() []inputs.Measurement {
	return getCollectorMeasurement()
}

func (i *Input) ElectionEnabled() bool { return i.Election }

func (i *Input) Terminate() {
	if i.semStop != nil {
		i.semStop.Close()
	}
}

func (i *Input) Pause() error {
	tick := time.NewTicker(inputs.ElectionPauseTimeout)
	select {
	case i.chPause <- true:
		return nil
	case <-tick.C:
		return fmt.Errorf("pause %s failed", inputName)
	}
}

func (i *Input) Resume() error {
	tick := time.NewTicker(inputs.ElectionResumeTimeout)
	select {
	case i.chPause <- false:
		return nil
	case <-tick.C:
		return fmt.Errorf("resume %s failed", inputName)
	}
}

func newInput() *Input {
	return &Input{
		Tags:                      make(map[string]string),
		LoggingExtraSourceMap:     make(map[string]string),
		LoggingSourceMultilineMap: make(map[string]string),
		Election:                  true,
		Feeder:                    dkio.DefaultFeeder(),
		Tagger:                    dkpt.DefaultGlobalTagger(),
		pause:                     &atomic.Bool{},
		chPause:                   make(chan bool, inputs.ElectionPauseChannelLength),
		semStop:                   cliutils.NewSem(),
	}
}

//nolint:gochecknoinits
func init() {
	inputs.Add(inputName, func() inputs.Input {
		return newInput()
	})
}
