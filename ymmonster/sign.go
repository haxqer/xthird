package ymmonster

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
	"sort"
	"strings"
)

func Sign(bm xthird.BodyMap, key string) string {
	bm.Remove("signature")
	//s := bm.EncodeCommonSignParams()
	s := encodeCommonSignParams(bm)
	s += "&key=" + key
	return gofunc.Sha1Lower(s)
}

func VerifySign(bm xthird.BodyMap, key string) bool {
	signature := bm.GetString("signature")
	return signature == Sign(bm, key)
}


// ("bar=baz&foo=quux") sorted by key.
func encodeCommonSignParams(bm xthird.BodyMap) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		 v := bm.GetString(k)
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
		buf.WriteByte('&')

	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}