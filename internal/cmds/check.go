// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package cmds

import (
	"fmt"
	"os"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	cp "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/colorprint"
)

func runCheckFlags() error {
	switch {
	case *flagCheckConfig:
		confdir := *flagCheckConfigDir
		if confdir == "" {
			tryLoadMainCfg()
			confdir = datakit.ConfdDir
		}

		if err := checkConfig(confdir, ".conf"); err != nil {
			os.Exit(-1)
		}
		os.Exit(0)

	case *flagCheckSNMP != "":
		if !datakit.FileExist(*flagCheckSNMP) {
			cp.Errorf("[E] File not exist: %s\n", *flagCheckSNMP)
			return nil
		}

		if err := testSNMP(*flagCheckSNMP); err != nil {
			os.Exit(-1)
		}
		os.Exit(0)

	case *flagCheckSample:
		if err := checkSample(); err != nil {
			os.Exit(-1)
		}
		os.Exit(0)
	}

	return fmt.Errorf("unknown check option: %s", os.Args[2])
}
