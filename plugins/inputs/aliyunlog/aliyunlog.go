package aliyunlog

import (
	"context"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/models"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	consumerLibrary "github.com/aliyun/aliyun-log-go-sdk/consumer"
)

type AliyunLog struct {
	Consumer []*ConsumerInstance

	runningInstances []*runningInstance

	ctx       context.Context
	cancelFun context.CancelFunc

	accumulator telegraf.Accumulator

	logger *models.Logger
}

type runningInstance struct {
	cfg *ConsumerInstance

	agent *AliyunLog

	logger *models.Logger

	runningProjects []*runningProject
}

type runningProject struct {
	inst *runningInstance
	cfg  *LogProject

	logger *models.Logger

	runningStores []*runningStore
}

type runningStore struct {
	proj       *runningProject
	cfg        *LogStoreCfg
	metricName string

	logger *models.Logger
}

func (_ *AliyunLog) Catalog() string {
	return "aliyun"
}

func (_ *AliyunLog) SampleConfig() string {
	return aliyunlogConfigSample
}

func (_ *AliyunLog) Description() string {
	return ""
}

func (_ *AliyunLog) Gather(telegraf.Accumulator) error {
	return nil
}

func (al *AliyunLog) Start(acc telegraf.Accumulator) error {

	al.logger = &models.Logger{
		Name: `aliyunlog`,
	}

	if len(al.Consumer) == 0 {
		al.logger.Warnf("no configuration found")
		return nil
	}

	al.logger.Info("starting...")

	al.accumulator = acc

	for _, instCfg := range al.Consumer {
		r := &runningInstance{
			cfg:    instCfg,
			agent:  al,
			logger: al.logger,
		}
		al.runningInstances = append(al.runningInstances, r)

		go r.run(al.ctx)
	}

	return nil
}

func (al *AliyunLog) Stop() {
	al.cancelFun()
}

func (r *runningInstance) run(ctx context.Context) error {

	for _, c := range r.cfg.Projects {

		p := &runningProject{
			cfg:    c,
			inst:   r,
			logger: r.logger,
		}
		r.runningProjects = append(r.runningProjects, p)

		go p.run(ctx)
	}

	return nil
}

func (r *runningProject) run(ctx context.Context) error {

	for _, c := range r.cfg.Stores {

		s := &runningStore{
			cfg:    c,
			proj:   r,
			logger: r.logger,
		}
		s.metricName = c.MetricName
		if s.metricName == "" {
			s.metricName = `aliyunlog_` + c.Name
		}
		r.runningStores = append(r.runningStores, s)

		go s.run(ctx)
	}

	return nil
}

func (r *runningStore) run(ctx context.Context) error {

	option := consumerLibrary.LogHubConfig{
		Endpoint:          r.proj.inst.cfg.Endpoint,
		AccessKeyID:       r.proj.inst.cfg.AccessKey,
		AccessKeySecret:   r.proj.inst.cfg.AccessID,
		Project:           r.proj.cfg.Name,
		Logstore:          r.cfg.Name,
		ConsumerGroupName: r.cfg.ConsumerGroupName,
		ConsumerName:      r.cfg.ConsumerName,
		// This options is used for initialization, will be ignored once consumer group is created and each shard has been started to be consumed.
		// Could be "begin", "end", "specific time format in time stamp", it's log receiving time.
		CursorPosition: consumerLibrary.BEGIN_CURSOR,
	}

	consumerWorker := consumerLibrary.InitConsumerWorker(option, r.logProcess)
	consumerWorker.Start()

	select {
	case <-ctx.Done():
		consumerWorker.StopAndWait()
	}

	r.logger.Infof("%s done", r.cfg.Name)

	return nil

}

func (r *runningStore) logProcess(shardId int, logGroupList *sls.LogGroupList) string {
	r.logger.Debugf("shardId:%d, grouplist:%s", shardId, logGroupList.String())
	for _, lg := range logGroupList.LogGroups {

		tags := map[string]string{}
		tags["store"] = r.cfg.Name
		tags["project"] = r.proj.cfg.Name

		for _, lt := range lg.GetLogTags() {
			k := lt.GetKey()
			if k == "" {
				continue
			}
			tags[k] = lt.GetValue()
		}

		if lg.GetSource() != "" {
			tags["source"] = lg.GetSource()
		}

		if lg.GetTopic() != "" {
			tags["topic"] = lg.GetTopic()
		}

		for _, l := range lg.GetLogs() {

			fields := map[string]interface{}{}

			for _, lc := range l.Contents {
				k := lc.GetKey()
				if k != "" {
					fields[k] = lc.GetValue()
				}
			}

			tm := time.Unix(int64(l.GetTime()), 0)
			m, err := metric.New(r.metricName, tags, fields, tm)
			if err == nil {
				r.proj.inst.agent.accumulator.AddMetric(m)
			} else {
				r.logger.Warnf("%s", err)
			}
		}
	}
	return ""
}

func init() {
	inputs.Add("aliyunlog", func() inputs.Input {
		ac := &AliyunLog{}
		ac.ctx, ac.cancelFun = context.WithCancel(context.Background())
		return ac
	})
}
