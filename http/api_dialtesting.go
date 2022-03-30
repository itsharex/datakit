package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	uhttp "gitlab.jiagouyun.com/cloudcare-tools/cliutils/network/http"
	dt "gitlab.jiagouyun.com/cloudcare-tools/kodo/dialtesting"
)

type dialtestingDebugRequest struct {
	Task  interface{} `json:"task"`
	Class string      `json:"class"`
}

type dialtestingDebugResponse struct {
	Cost         string `json:"cost"`
	ErrorMessage string `json:"error_msg"`
	Status       string `json:"status"`
	Traceroute   string `json:"traceroute"`
}

func apiDebugDialtestingHandler(w http.ResponseWriter, req *http.Request, whatever ...interface{}) (interface{}, error) {
	var (
		tid        = req.Header.Get(uhttp.XTraceId)
		start      = time.Now()
		t          dt.Task
		traceroute string
		status     = "success"
	)

	reqDebug, err := getAPIDebugDialtestingRequest(req)
	if err != nil {
		l.Errorf("[%s] %s", tid, err.Error())
		return nil, uhttp.Error(ErrInvalidRequest, err.Error())
	}

	switch reqDebug.Class {
	case dt.ClassHTTP:
		t = &dt.HTTPTask{}
	case dt.ClassTCP:
		t = &dt.TcpTask{}
	case dt.ClassWebsocket:
		t = &dt.WebsocketTask{}
	case dt.ClassICMP:
		t = &dt.IcmpTask{}
	default:
		l.Errorf("unknown task type: %s", reqDebug.Class)
		return nil, uhttp.Error(ErrInvalidRequest, fmt.Sprintf("unknown task type:%s", reqDebug.Class))
	}

	bys, err := json.Marshal(reqDebug.Task)
	if err != nil {
		l.Errorf(`json.Marshal: %s`, err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bys, &t); err != nil {
		l.Errorf(`json.Unmarshal: %s`, err.Error())
		return nil, err
	}
	if strings.ToLower(t.Status()) == dt.StatusStop {
		return nil, uhttp.Error(ErrInvalidRequest, "the task status is stop")
	}

	// -- dialtesting debug procedure start --
	if err := defDialtestingMock.debugInit(t); err != nil {
		l.Errorf("[%s] %s", tid, err.Error())
		return nil, uhttp.Error(ErrInvalidRequest, err.Error())
	}
	if err := defDialtestingMock.debugRun(t); err != nil {
		l.Errorf("[%s] %s", tid, err.Error())
		return nil, uhttp.Error(ErrInvalidRequest, err.Error())
	}

	_, fields := defDialtestingMock.getResults(t)

	failReason, ok := fields["fail_reason"].(string)
	if ok {
		status = "fail"
	}
	if reqDebug.Class == dt.ClassTCP || reqDebug.Class == dt.ClassICMP {
		traceroute, _ = fields["traceroute"].(string)
	}
	if reqDebug.Class == dt.ClassTCP {
		responseTime, _ := fields["response_time"].(int64)
		if responseTime == 0 {
			status = "timeout"
		}
	}
	if reqDebug.Class == dt.ClassICMP {
		lossPercent, _ := fields["packet_loss_percent"].(float64)
		if lossPercent == 100 {
			status = "timeout"
		}
	}

	return &dialtestingDebugResponse{
		Cost:         time.Since(start).String(),
		ErrorMessage: failReason,
		Status:       status,
		Traceroute:   traceroute,
	}, nil
}

func getAPIDebugDialtestingRequest(req *http.Request) (*dialtestingDebugRequest, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, uhttp.Error(ErrInvalidRequest, err.Error())
	}

	var reqDebug dialtestingDebugRequest
	if err := json.Unmarshal(body, &reqDebug); err != nil {
		return nil, uhttp.Error(ErrInvalidRequest, err.Error())
	}

	return &reqDebug, nil
}
