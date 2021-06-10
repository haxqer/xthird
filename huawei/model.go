package huawei

type TokenAuthResponse struct {
	ClientId string `json:"client_id"` // 应用id
	ExpireIn int    `json:"expire_in,omitempty"`
	UnionId  string `json:"union_id"`
	OpenId   string `json:"open_id"`
	Scope    string `json:"scope"`
}

type GetUserInfoResponse struct {
	OpenID         string `json:"openID"` // 用户openID
	DisplayName    string `json:"displayName"`
	HeadPictureURL string `json:"headPictureURL,omitempty"` // 头像
	MobileNumber   string `json:"mobileNumber,omitempty"`   // 手机号
	BirthDate      string `json:"birthDate,omitempty"`      // 生日
	AgeGroupFlag   int    `json:"ageGroupFlag,omitempty"`   // 年龄段
	Email          string `json:"email,omitempty"`          // 年龄段
}
