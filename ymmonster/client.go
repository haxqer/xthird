package ymmonster

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
