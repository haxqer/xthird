package ymmonster

import (
	"github.com/haxqer/xthird"
	"net/http"
	"net/url"
)

func ParseNotifyByQueryURLValues(req *http.Request) (bm xthird.BodyMap, err error) {
	return ParseNotifyByURLValues(req.URL.Query())
}

func ParseNotifyByURLValues(value url.Values) (bm xthird.BodyMap, err error) {
	bm = make(xthird.BodyMap, len(value)+1)
	for k, v := range value {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return
}
