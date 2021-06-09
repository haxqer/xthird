package yyb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
	"net/url"
)

func (c *Client) AuthToken(bm xthird.BodyMap, loginType string) (rsp *TokenAuthResponse, err error) {
	err = bm.CheckEmptyError("openid", "openkey")
	if err != nil {
		return nil, err
	}
	if loginType != QQ && loginType != WECHAT {
		return nil, errors.New("loginType error")
	}
	bs, err := c.doAuthToken(bm, loginType)
	if err != nil {
		return nil, err
	}
	rsp = new(TokenAuthResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return rsp, nil

}

func (c *Client) doAuthToken(bm xthird.BodyMap, loginType string) (bs []byte, err error) {
	var (
		reqUrl string
		appKey string
		appId  string
	)

	if loginType == WECHAT {
		reqUrl = AuthTokenWechatTestUrl
		appId = c.WechatAppId
		appKey = c.WechatAppSecret
	} else if loginType == QQ {
		appKey = c.SanBoxAppKey
		reqUrl = AuthTokenQQTestUrl
		appId = c.QQAppId
	}

	bm.Set("sig", SignAuthToken(bm, appKey))
	bm.Set("appid", appId)
	param := FormatURLParam(bm)
	httpClient := xhttp.NewClient()
	res, bs, errs := httpClient.Type(xhttp.TypeUrlencoded).Get(reqUrl + "?" + param).EndBytes()
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
