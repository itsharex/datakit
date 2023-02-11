// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"testing"
	"time"

	tu "github.com/GuanceCloud/cliutils/testutil"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/ptinput"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type funcCase struct {
	name     string
	in       string
	script   string
	expected interface{}
	key      string
}

func TestDecode(t *testing.T) {
	data := []string{"测试一下", "不知道", "测试一下123456", "哈哈哈哈哈", "-汪98阿萨德离开家"}
	decode_data_slice := make([]string, 10)

	for idx, cont := range data {
		decode_data, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(cont))
		decode_data_slice[idx] = string(decode_data)
	}

	testCase := []*funcCase{
		{
			in:     decode_data_slice[0],
			script: `decode(_,"gbk")`,
			key:    "message",
		},
		{
			in:     decode_data_slice[1],
			script: `decode(_,"gbk")`,
			key:    "message",
		},
		{
			in:     decode_data_slice[2],
			script: `decode(_,"gbk")`,
			key:    "message",
		},
		{
			in:     decode_data_slice[3],
			script: `decode(_,"gbk")`,
			key:    "message",
		},
		{
			in:     decode_data_slice[4],
			script: `decode(_,"gbk")`,
			key:    "message",
		},
	}
	for idx, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			runner, err := NewTestingRunner(tc.script)
			tu.Equals(t, nil, err)

			pt := ptinput.GetPoint()
			ptinput.InitPt(pt, "test", nil, map[string]any{"message": tc.in}, time.Now())
			errR := runScript(runner, pt)

			if errR != nil {
				ptinput.PutPoint(pt)
				t.Fatal(errR)
			}

			tu.Equals(t, data[idx], pt.Fields[tc.key])

			t.Logf("[%d] PASS", idx)
			ptinput.PutPoint(pt)
		})
	}
}
