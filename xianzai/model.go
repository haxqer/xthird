package xianzai

type OrderResponse struct {
	ResponseCode string `schema:"responseCode"`
	PayUrl       string `schema:"tn"`
}
