package ymmonster

type AuthResponse struct {
	Code       string    `json:"code"`
	Msg        string `json:"msg,omitempty"`
}
