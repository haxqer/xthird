package yyb

type TokenAuthResponse struct {
	ReturnCode int    `json:"ret"` // ret 为 0 则为成功
	Message    string `json:"msg,omitempty"`
}

type OrderResponse struct {
	ReturnCode int    `json:"ret"` // ret 为 0 则为成功
	Message    string `json:"msg,omitempty"`
}
