package meizu

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
)

func Sign(bm xthird.BodyMap, appSecret string) string {
	bm.Remove("sign_type")
	bm.Remove("sign")
	s := bm.EncodeCommonSignParamsAllowEmpty() + ":" + appSecret
	return gofunc.Md5Lower(s)
}
