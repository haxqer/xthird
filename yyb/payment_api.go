package yyb

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
	"hash"
	"net/url"
	"time"
)

func (c *Client) Order(bm xthird.BodyMap, loginType string) (rsp *OrderResponse, err error) {
	err = bm.CheckEmptyError("openid", "openkey", "pf", "pfkey", "payitem", "goodsmeta", "zoneid")
	if err != nil {
		return nil, err
	}
	if loginType != QQ && loginType != WECHAT {
		return nil, errors.New("loginType error")
	}
	bs, err := c.doOrder(bm, loginType)
	if err != nil {
		return nil, err
	}
	rsp = new(OrderResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return rsp, nil
}

func (c *Client) doOrder(bm xthird.BodyMap, loginType string) (bs []byte, err error) {
	var (
		reqUrl       string
		appKey       string
		appId        string
		ySessionId   string
		ySessionType string
		yOrgLoc      string
	)
	reqUrl = OrderTestUrl
	yOrgLoc = "/v3/r/mpay/buy_goods_m"

	if loginType == WECHAT {
		appId = c.WechatAppId
		appKey = c.WechatAppSecret
		ySessionId = "hy_gameid"
		ySessionType = "wc_actoken"
	} else if loginType == QQ {
		appKey = c.SanBoxAppKey
		appId = c.QQAppId
		ySessionId = "openid"
		ySessionType = "kp_actoken"
	}

	bm.Set("appid", appId)
	bm.Set("ts", gofunc.Int64ToString(time.Now().Unix()))
	baseStr := "GET&" + url.QueryEscape(yOrgLoc) + "&" + url.QueryEscape(bm.EncodeYYBSignParams())
	var h hash.Hash
	h = hmac.New(sha1.New, []byte(appKey+"&"))
	h.Write([]byte(baseStr))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	bm.Set("sig", sign)
	param := FormatURLParam(bm)

	httpClient := xhttp.NewClient()
	cookie := fmt.Sprintf("session_id=%s;session_type=%s;org_loc=%s", ySessionId, ySessionType, yOrgLoc)
	httpClient.Header.Add("Cookie", cookie)
	reqUrl = reqUrl + "?" + param

	res, bs, errs := httpClient.Type(xhttp.TypeUrlencoded).Get(reqUrl).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}

	return bs, nil
}
