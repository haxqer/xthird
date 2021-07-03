package meizu

type AuthResponse struct {
	Code    int `json:"code"`
	Message string `json:"message,omitempty"`
	Value   string `json:"value,omitempty"`
}
