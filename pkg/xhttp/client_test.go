package xhttp

import (
	"testing"
	"time"

	"github.com/haxqer/xthird/pkg/xlog"
)

type HttpGet struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func TestHttpGet(t *testing.T) {
	client := NewClient()
	client.Timeout = 10 * time.Second

	rsp := new(HttpGet)
	_, errs := client.Type(TypeJSON).Get("https://wwww.baidu.com").EndStruct(rsp)
	if len(errs) > 0 {
		xlog.Error(errs[0])
		return
	}
	xlog.Debug(rsp)

	// test
	_, bs, errs := client.Get("http://www.baidu.com").EndBytes()
	if len(errs) > 0 {
		xlog.Error(errs[0])
		return
	}
	xlog.Debug(string(bs))
}
