// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package io

import (
	"fmt"
	"time"

	"github.com/GuanceCloud/cliutils/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io/dataway"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/io/failcache"
)

func (x *dkIO) cacheData(c *consumer, d *feedOption, tryClean bool) {
	if d == nil {
		log.Warn("get empty data, ignored")
		return
	}

	if len(d.pts) == 0 {
		log.Warnf("no point from %q", d.input)
		return
	}

	defer func() {
		queuePtsVec.WithLabelValues(d.cat.String()).Set(float64(len(c.points)))
	}()

	log.Debugf("get iodata(%d points) from %s|%s", len(d.pts), d.cat, d.input)

	x.recordPoints(d)
	c.points = append(c.points, d.pts...)

	if tryClean && x.maxCacheCount > 0 && len(c.points) > x.maxCacheCount {
		x.flush(c)

		// reset consumer flush ticker to prevent send small packages
		c.flushTiker.Reset(x.flushInterval)
	}
}

func (x *dkIO) recordPoints(d *feedOption) {
	if x.recorder != nil && x.recorder.Enabled {
		if err := x.recorder.Record(d.pts, d.cat, d.input); err != nil {
			log.Warnf("record %d points on %q from %q failed: %s", len(d.pts), d.cat, d.input, err)
		} else {
			log.Debugf("record %d points on %q from %q ok", len(d.pts), d.cat, d.input)
		}
	} else {
		log.Debugf("recorder disabled: %+#v", x.recorder)
	}
}

func (x *dkIO) flush(c *consumer) {
	c.lastFlush = time.Now()

	defer func() {
		flushVec.WithLabelValues(c.category.String()).Inc()
	}()

	if err := x.doFlush(c.points, c.category, c.fc); err != nil {
		log.Warnf("post %d points to %s failed: %s, ignored", len(c.points), c.category, err)
	}

	// I think here is the best position to put back these points.
	datakit.PutbackPoints(c.points...)

	c.points = c.points[:0] // clear
}

func (x *dkIO) flushFailCache(c *consumer) {
	if c.fc == nil {
		return
	}

	if err := x.dw.Write(dataway.WithCacheClean(true),
		dataway.WithCategory(c.category),
		dataway.WithFailCache(c.fc),
	); err != nil {
		log.Warnf("flush cache failed: %s, ignored", err)
	}
}

func (x *dkIO) doFlush(points []*point.Point,
	cat point.Category,
	fc failcache.Cache,
	dynamicURL ...string,
) error {
	if x.dw == nil {
		return fmt.Errorf("dataway not set")
	}

	if len(points) == 0 {
		return nil
	}

	opts := []dataway.WriteOption{
		dataway.WithPoints(points),

		// max cache size(in memory) upload as a batch
		dataway.WithBatchSize(x.maxCacheCount),

		dataway.WithCategory(cat),
		dataway.WithFailCache(fc),
		dataway.WithCacheAll(x.cacheAll),
	}

	if len(dynamicURL) > 0 {
		opts = append(opts, dataway.WithDynamicURL(dynamicURL[0]))
	}

	return x.dw.Write(opts...)
}
