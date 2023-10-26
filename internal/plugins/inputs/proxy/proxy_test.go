// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package proxy

import (
	"testing"

	"github.com/GuanceCloud/cliutils/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/testutils"
)

func TestProxyServer(t *testing.T) {
	var pts []*point.Point
	for i := 0; i < 100; i++ {
		pts = append(pts, testutils.RandPointV2("test_point", 10, 30))
	}

	for _, pt := range pts {
		log.Info(pt.MustLPPoint().String())
	}
}
