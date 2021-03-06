package xiaomi

import (
	"errors"
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
	"net/url"
	"sort"
	"strings"
)

func Sign(bm xthird.BodyMap, key string) string {
	bm.Remove("signature")
	s := bm.EncodeCommonSignParams()
	return gofunc.HmacSha1HexString(s, key)
}

func VerifySign(bm xthird.BodyMap, key string) bool {
	//err := bm.CheckEmptyError("appId", "cpOrderId", "uid", "orderId", "orderStatus", "payFee", "productCode", "productCode", "productCount", "signature")
	//if err != nil {
	//	return false
	//}
	signature := bm.GetString("signature")
	bm.Remove("signature")
	s, err := buildSignStr(bm)
	if err != nil {
		return false
	}
	return signature == gofunc.HmacSha1HexString(s, key)
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
			unescape, err := url.QueryUnescape(v)
			if err != nil {
				return xthird.NULL, err
			}
			buf.WriteString(unescape)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return xthird.NULL, errors.New("length is error")
	}
	return buf.String()[:buf.Len()-1], nil
}