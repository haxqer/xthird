package xianzai

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
)

func Sign(bm xthird.BodyMap, key string) string {
	bm.Remove("mhtSignature")
	bm.Remove("mhtSignType")
	s := bm.EncodeCommonSignParams()
	s += "&" + gofunc.Md5Lower(key)
	return gofunc.Md5Lower(s)
}


