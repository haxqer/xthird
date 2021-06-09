package vivo

type TokenAuthResponse struct {
	ReturnCode int                    `json:"retcode"`
	Data       *TokenAuthResponseData `json:"data,omitempty"`
}

type TokenAuthResponseData struct {
	Success bool   `json:"success,omitempty"`
	OpenId  string `json:"openid,omitempty"`
}
