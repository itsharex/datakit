// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package trace

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/lineproto"
	dkio "gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/point"
)

func TestAfterGather(t *testing.T) {
	dkioFeed = func(name, category string, pts []*point.Point, opt *dkio.Option) error { return nil }

	StartTracingStatistic()

	afterGather := NewAfterGather()
	afterGather.AppendCalculator(StatTracingInfo)

	afterGather.AppendFilter(PenetrateErrorTracing)

	closer := &CloseResource{}
	closer.UpdateIgnResList(map[string][]string{_services[0]: _resources})
	afterGather.AppendFilter(closer.Close)

	keeper := &KeepRareResource{}
	keeper.UpdateStatus(true, time.Second)
	afterGather.AppendFilter(keeper.Keep)

	sampler := &Sampler{SamplingRateGlobal: 0.33}
	afterGather.AppendFilter(sampler.Sample)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				trace := randDatakitTrace(t, 10, randService(_services...), randResource(_resources...))
				parentialize(trace)
				afterGather.Run("test_after_gather", DatakitTraces{trace}, false)
			}
		}()
	}
	wg.Wait()
}

func TestBuildPoint(t *testing.T) {
	encoder := lineproto.NewLineEncoder()
	for i := 0; i < 100; i++ {
		pt, err := BuildPoint(randDatakitSpan(t), false)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
		err = encoder.AppendPoint(pt.Point)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	lines, err := encoder.UnsafeString()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lines)
}

func TestBuildPointsBatch(t *testing.T) {
	for i := 0; i < 100; i++ {
		BuildPointsBatch(DatakitTraces{randDatakitTrace(t, 10)}, false)
	}
}
