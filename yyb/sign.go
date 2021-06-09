package yyb

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
)

func SignAuthToken(bm xthird.BodyMap, appKey string) string {
	return gofunc.Md5Lower(appKey + bm.GetString("timestamp"))
}

func SignOrder(bm xthird.BodyMap, appKey, uri string)  {

}