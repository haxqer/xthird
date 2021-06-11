package ymmonster

import (
	"encoding/json"
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
	"time"
)

func(x Client) AuthToken(bm xthird.BodyMap) (rsp *AuthResponse, err error) {
	err = bm.CheckEmptyError("userId", "token")
	if err != nil {
		return nil, err
	}
	bs, err := x.doAuthToken(bm)
	if err != nil {
		return nil, err
	}
	rsp = new(AuthResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return rsp, nil
}

func (x Client) doAuthToken(bm xthird.BodyMap) (bs []byte, err error) {

	bm.Set("appId", x.AppId)
	bm.Set("timestamp", time.Now().Unix())


	signature := Sign(bm, x.AppKey)
	bm.Set("signature", signature)
	param := bm.EncodeGetParams()

	httpClient := xhttp.NewClient()

	//return nil, nil
	res, bs, errs := httpClient.Type(xhttp.TypeFormData).Post("http://sdk.plat.pkey.cn/web-mapi/api/checkToken.do").SendString(param).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil

}

