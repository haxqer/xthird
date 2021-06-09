package xiaomi

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
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
	return bm.GetString("signature") == Sign(bm, key)
}