package xiaomi

import (
	"encoding/json"
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
)

func (c *Client) AuthToken(bm xthird.BodyMap) (rsp *AuthResponse, err error) {
	err = bm.CheckEmptyError("appId", "session", "uid")
	if err != nil {
		return nil, err
	}
	bs, err := c.doAuthToken(bm)
	if err != nil {
		return nil, err
	}
	rsp = new(AuthResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return rsp, nil
}

func (c *Client) doAuthToken(bm xthird.BodyMap) (bs []byte, err error) {
	sign := Sign(bm, c.AppSecret)
	bm.Set("signature", sign)
	param := bm.FormatURLParam()
	httpClient := xhttp.NewClient()
	res, bs, errs := httpClient.Type(xhttp.TypeForm).Post(AuthTokenUrl).SendString(param).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}
