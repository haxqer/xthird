package huawei

import (
	"encoding/json"
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
)

// https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides-V5/order-verify-purchase-token-0000001050033078-V5
// https://developer.huawei.com/consumer/cn/doc/development/HMSCore-References-V5/api-order-verify-purchase-token-0000001050746113-V5
// https://developer.huawei.com/consumer/cn/doc/development/HMSCore-References-V5/server-data-model-0000001050986133-V5#ZH-CN_TOPIC_0000001050986133__section264617465219

func (c *Client) OrderVerify(bm xthird.BodyMap) (rsp *OrderResponse, err error) {
	err = bm.CheckEmptyError( "purchaseToken", "productId")
	if err != nil {
		return nil, err
	}
	bs, err := c.doOrderPost(bm, OrderUrl)
	if err != nil {
		return nil, err
	}

	rspTemp := new(TempOrderResponse)
	if err = json.Unmarshal(bs, rspTemp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	println(rspTemp.PurchaseTokenData)
	td := new(PurchaseTokenData)
	if err = json.Unmarshal([]byte(rspTemp.PurchaseTokenData), td); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", rspTemp.PurchaseTokenData, err)
	}

	return &OrderResponse{
		ResponseCode:       rspTemp.ResponseCode,
		SignatureAlgorithm: rspTemp.SignatureAlgorithm,
		PurchaseTokenData:  td,
	}, nil
}


func (c *Client) doOrderPost(bm xthird.BodyMap, url string) (bs []byte, err error)  {
	httpClient := xhttp.NewClient()
	authHeaderString, err := c.BuildAuthorization()
	if err != nil {
		return nil, err
	}
	httpClient.Header.Add("Authorization", authHeaderString)

	res, bs, errs := httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}



