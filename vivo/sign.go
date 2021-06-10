package vivo

import (
	"errors"
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
	"sort"
	"strings"
)

func VerifySign(bm xthird.BodyMap, key string) bool {
	signature := bm.GetString("signature")
	bm.Remove("signature")
	bm.Remove("signMethod")
	return signature == sign(bm, key)
}

func sign(bm xthird.BodyMap, key string) string {
	s, _ := buildSignStr(bm)
	s += "&" + gofunc.Md5Lower(key)
	return gofunc.Md5Lower(s)
}

func buildSignStr(bm xthird.BodyMap) (string, error) {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := bm.GetString(k); v != xthird.NULL {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return xthird.NULL, errors.New("length is error")
	}
	return buf.String()[:buf.Len()-1], nil
}
