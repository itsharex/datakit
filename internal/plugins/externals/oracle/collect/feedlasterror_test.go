// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package collect

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/GuanceCloud/cliutils/logger"
)

func TestFeedLastErrorLoop(t *testing.T) {
	const defaultErrInterval = time.Second * 30

	l = logger.SLogger(InputName)
	datakitLastErrURL = "http://127.0.0.1:9529/v1/lasterror"

	defer func() {
		// Restore variables' value.
		l = nil
		datakitLastErrURL = ""
	}()

	type args struct {
		errString string
		ch        chan os.Signal
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				errString: "something",
				ch:        make(chan os.Signal),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				feedErrInterval = time.Second * 3
				defer func() {
					// Restore default values.
					feedErrInterval = defaultErrInterval
				}()

				timeout := time.NewTicker(time.Second * 10)
				defer timeout.Stop()

				for range timeout.C {
					tt.args.ch <- os.Interrupt
					break
				}
			}()

			FeedLastErrorLoop(tt.args.errString, tt.args.ch)
		})
	}
}

func TestPrint(t *testing.T) {
	err := fmt.Errorf("some error")

	t.Run("Println", func(t *testing.T) {
		fmt.Println("failed:", err.Error())
	})
}
