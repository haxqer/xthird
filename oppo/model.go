package oppo

type LoginResponse struct {
	ResultCode   string `xml:"resultCode,omitempty" json:"resultCode,omitempty"`
	ResultMsg    string `xml:"resultMsg,omitempty" json:"resultMsg,omitempty"`
	SsoId        string `xml:"ssoid,omitempty" json:"ssoid,omitempty"`
	UserName     string `xml:"userName,omitempty" json:"userName,omitempty"`
	Email        string `xml:"email,omitempty" json:"email,omitempty"`
	MobileNumber string `xml:"mobileNumber,omitempty" json:"mobileNumber,omitempty"`
}
