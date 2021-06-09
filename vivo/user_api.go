package vivo

import (
	"encoding/json"
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
	"net/url"
)

func AuthToken(bm xthird.BodyMap) (rsp *TokenAuthResponse, err error)  {
	err = bm.CheckEmptyError("opentoken")
	if err != nil {
		return nil, err
	}
	bs, err := doAuthToken(bm)
	if err != nil {
		return nil, err
	}
	rsp = new(TokenAuthResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return rsp, nil
}

func doAuthToken(bm xthird.BodyMap) (bs []byte, err error)  {
	param := FormatURLParam(bm)
	httpClient := xhttp.NewClient()
	res, bs, errs := httpClient.Type(xhttp.TypeFormData).Post(AuthTokenUrl).SendString(param).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// 格式化请求URL参数
func FormatURLParam(body xthird.BodyMap) (urlParam string) {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	return v.Encode()
}
