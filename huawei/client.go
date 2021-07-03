package huawei

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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

type AtResponse struct {
	AccessToken string `json:"access_token"`
}

var RequestHttpClient = http.Client{Timeout: time.Second * 10}

func (c *Client) GetAppAt() (string, error) {
	urlValue := url.Values{"grant_type": {"client_credentials"}, "client_secret": {c.AppSecret}, "client_id": {c.AppId}}
	resp, err := RequestHttpClient.PostForm("https://oauth-login.cloud.huawei.com/oauth2/v2/token", urlValue)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var atResponse AtResponse
	json.Unmarshal(bodyBytes, &atResponse)
	if atResponse.AccessToken != "" {
		return atResponse.AccessToken, nil
	} else {
		return "", errors.New("Get token fail, " + string(bodyBytes))
	}
}

func (c *Client) BuildAuthorization() (string, error) {
	appAt, err := c.GetAppAt()
	if err != nil {
		return "", err
	}
	oriString := fmt.Sprintf("APPAT:%s", appAt)
	var authString = base64.StdEncoding.EncodeToString([]byte(oriString))
	var authHeaderString = fmt.Sprintf("Basic %s", authString)
	return authHeaderString, nil
}