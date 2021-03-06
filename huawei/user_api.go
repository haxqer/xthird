package huawei

import (
	"encoding/json"
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
)

// https://developer.huawei.com/consumer/cn/doc/development/HMSCore-References-V5/account-gettokeninfo-0000001050050585-V5

func AuthToken(bm xthird.BodyMap) (rsp *AuthResponse, err error)  {
	err = bm.CheckEmptyError("access_token")
	if err != nil {
		return nil, err
	}
	bm.Set("open_id", "OPENID")
	bs, err := doUserPost(bm, AuthTokenUrl)
	if err != nil {
		return nil, err
	}
	rsp = new(AuthResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return rsp, nil
}

func doUserPost(bm xthird.BodyMap, url string) (bs []byte, err error)  {
	param := bm.FormatURLParam()
	httpClient := xhttp.NewClient()
	res, bs, errs := httpClient.Type(xhttp.TypeFormData).Post(url).SendString(param).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

