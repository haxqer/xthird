package ymmonster

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
)

func Sign(bm xthird.BodyMap, key string) string {
	bm.Remove("signature")
	s := bm.EncodeCommonSignParams()
	s += "&key=" + key
	return gofunc.Sha1Lower(s)
}

