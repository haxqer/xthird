package xthird

import (
	"net/http"
	"net/url"
)

func ParseNotifyToBodyMap(req *http.Request) (bm BodyMap, err error) {
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	var form map[string][]string = req.Form
	bm = make(BodyMap, len(form)+1)
	for k, v := range form {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return
}

func ParseNotifyByURLValues(value url.Values) (bm BodyMap, err error) {
	bm = make(BodyMap, len(value)+1)
	for k, v := range value {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return
}