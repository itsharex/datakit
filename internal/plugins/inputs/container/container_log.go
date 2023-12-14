// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package container

import (
	"context"
	"fmt"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/container/runtime"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/goroutine"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/plugins/inputs"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/tailer"
)

func (c *container) cleanMissingContainerLog(newIDs []string) {
	missingIDs := c.logTable.findDifferences(newIDs)
	for _, id := range missingIDs {
		l.Infof("clean log collection for container id %s", id)
		c.logTable.closeFromTable(id)
		c.logTable.removeFromTable(id)
	}
}

func (c *container) tailingLogs(ins *logInstance) {
	g := goroutine.NewGroup(goroutine.Option{Name: "container-logs/" + ins.containerName})
	done := make(chan interface{})

	for _, cfg := range ins.configs {
		if cfg.Disable {
			continue
		}

		path := cfg.Path
		if cfg.HostFilePath != "" {
			path = cfg.HostFilePath
			l.Infof("container log %s redirect to host path %s", cfg.Path, cfg.HostFilePath)
		}

		mergedTags := inputs.MergeTags(c.extraTags, cfg.Tags, "")

		opt := &tailer.Option{
			Source:                         cfg.Source,
			Service:                        cfg.Service,
			Pipeline:                       cfg.Pipeline,
			CharacterEncoding:              cfg.CharacterEncoding,
			MultilinePatterns:              cfg.MultilinePatterns,
			GlobalTags:                     mergedTags,
			BlockingMode:                   c.ipt.LoggingBlockingMode,
			MaxMultilineLifeDuration:       c.ipt.LoggingMaxMultilineLifeDuration,
			RemoveAnsiEscapeCodes:          c.ipt.LoggingRemoveAnsiEscapeCodes,
			MaxForceFlushLimit:             c.ipt.LoggingForceFlushLimit,
			FileFromBeginningThresholdSize: int64(c.ipt.LoggingFileFromBeginningThresholdSize),
			Done:                           done,
		}

		switch cfg.Type {
		case "file":
			opt.Mode = tailer.FileMode
		case runtime.DockerRuntime:
			opt.Mode = tailer.DockerMode
		default:
			opt.Mode = tailer.ContainerdMode
		}

		_ = opt.Init()

		path = logsJoinRootfs(path)

		filelist, err := tailer.NewProvider().SearchFiles([]string{path}).Result()
		if err != nil {
			l.Warnf("failed to scan container-log collection %s(%s) for %s, err: %s", cfg.Path, path, ins.containerName, err)
			continue
		}

		if len(filelist) == 0 {
			l.Infof("container %s not found any log file for path %s, skip", ins.containerName, path)
			continue
		}

		for _, file := range filelist {
			if c.logTable.inTable(ins.id, file) {
				continue
			}

			l.Infof("add container log collection with path %s from source %s", file, opt.Source)

			tail, err := tailer.NewTailerSingle(file, opt)
			if err != nil {
				l.Errorf("failed to create container-log collection %s for %s, err: %s", file, ins.containerName, err)
				continue
			}

			c.logTable.addToTable(ins.id, file, done)

			func(file string) {
				g.Go(func(ctx context.Context) error {
					defer func() {
						c.logTable.removePathFromTable(ins.id, file)
						l.Infof("remove container log collection from source %s", opt.Source)
					}()
					tail.Run()
					return nil
				})
			}(file)
		}
	}
}

func (c *container) queryContainerLogInfo(info *runtime.Container) *logInstance {
	podName := getPodNameForLabels(info.Labels)
	podNamespace := getPodNamespaceForLabels(info.Labels)

	ins := &logInstance{
		id:            info.ID,
		containerName: info.Name,
		image:         info.Image,
		logPath:       info.LogPath,
		podName:       podName,
		podNamespace:  podNamespace,
		volMounts:     info.Mounts,
	}

	if name := getContainerNameForLabels(info.Labels); name != "" {
		ins.containerName = name
	}

	if c.k8sClient != nil && podName != "" {
		podInfo, err := c.queryPodInfo(context.Background(), podName, podNamespace)
		if err != nil {
			l.Warn(err)
		} else {
			// ex: datakit/logs
			if v := podInfo.pod.Annotations[fmt.Sprintf(logConfigAnnotationKeyFormat, "")]; v != "" {
				ins.configStr = v
			}

			// ex: datakit/nginx.logs
			if v := podInfo.pod.Annotations[fmt.Sprintf(logConfigAnnotationKeyFormat, ins.containerName+".")]; v != "" {
				ins.configStr = v
			}

			ins.podLabels = podInfo.pod.Labels
			ins.ownerKind, ins.ownerName = podInfo.owner()

			// use Image from Pod Container
			if img := podInfo.containerImage(ins.containerName); img != "" {
				ins.image = img
			}
		}
	}

	// ex: DATAKIT_LOGS_CONFIG
	if info.Envs != nil {
		if str, ok := info.Envs["DATAKIT_LOGS_CONFIG"]; ok {
			ins.configStr = str
		}
	}

	ins.imageName, ins.imageShortName, ins.imageTag = runtime.ParseImage(ins.image)

	l.Debugf("container %s use config: %v", ins.containerName, ins)
	return ins
}
