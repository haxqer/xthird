package huawei

type TempOrderResponse struct {
	ResponseCode       string `json:"responseCode"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	PurchaseTokenData  string `json:"purchaseTokenData"`
}

type OrderResponse struct {
	ResponseCode       string             `json:"responseCode"`
	SignatureAlgorithm string             `json:"signatureAlgorithm"`
	PurchaseTokenData  *PurchaseTokenData `json:"purchaseTokenData"`
}

type PurchaseTokenData struct {
	ApplicationId    int64  `json:"applicationId"`
	OrderId          string `json:"orderId"`
	Kind             int    `json:"kind"`
	ProductId        string `json:"productId"`
	PurchaseState    int    `json:"purchaseState"`
	Price            int    `json:"price"`
	DeveloperPayload string `json:"developerPayload"`
}

type AuthResponse struct {
	ClientId  string `json:"client_id"` // 应用id
	ExpireIn  int    `json:"expire_in,omitempty"`
	UnionId   string `json:"union_id"`
	OpenId    string `json:"open_id"`
	Scope     string `json:"scope"`
	ProjectId string `json:"project_id"`
}

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
