package xianzai

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
)

func Sign(bm xthird.BodyMap, key string) string {
	bm.Remove("mhtSignature")
	bm.Remove("signature")
	s := bm.EncodeCommonSignParams() + "&" + gofunc.Md5Lower(key)
	return gofunc.Md5Lower(s)
}

func VerifySign(bm xthird.BodyMap, key string)  {

}