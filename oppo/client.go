package oppo

import (
	"encoding/json"
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
	"net/url"
)

type Client struct {
	AppId     string
	AppKey    string
	AppSecret string
}

func NewClient(appId, appKey, appSecret string) (client *Client) {
	return &Client{
		AppId:     appId,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (o *Client) Login(bm xthird.BodyMap) (oppoRsp *LoginResponse, err error) {
	err = bm.CheckEmptyError("oss_id", "token")
	if err != nil {
		return nil, err
	}
	bs, err := o.doGet(bm)
	if err != nil {
		return nil, err
	}
	oppoRsp = new(LoginResponse)
	if err = json.Unmarshal(bs, oppoRsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return oppoRsp, nil
}

func (o *Client) doGet(bm xthird.BodyMap) (bs []byte, err error) {

	baseStr, sign := GenLoginBaseStr(bm, o.AppKey, o.AppSecret)

	reqUrl := fmt.Sprintf(LoginUrl, bm["oss_id"], url.QueryEscape(bm["token"].(string)))

	httpClient := xhttp.NewClient()
	httpClient.Header.Add("param", baseStr)
	httpClient.Header.Add("oauthSignature", sign)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Get(reqUrl).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}

	return bs, nil
}
