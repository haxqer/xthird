package xiaomi

type AuthResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errMsg,omitempty"`
	Adult   int    `json:"adult,omitempty"`
	Age     int    `json:"age,omitempty"`
}
