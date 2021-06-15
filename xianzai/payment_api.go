package xianzai

import (
	"fmt"
	"github.com/haxqer/xthird"
	"github.com/haxqer/xthird/pkg/xhttp"
	"net/url"
	"time"
)

func (x Client) Order(bm xthird.BodyMap) (rsp *OrderResponse, err error) {
	err = bm.CheckEmptyError("frontNotifyUrl", "mhtOrderAmt", "mhtOrderDetail", "mhtOrderName", "mhtOrderNo", "notifyUrl", "consumerCreateIp")
	if err != nil {
		return nil, err
	}
	bs, err := x.doOrder(bm)
	if err != nil {
		return nil, err
	}
	m, err := url.ParseQuery(string(bs))
	if err != nil {
		return nil, fmt.Errorf("url.ParseQuery(%s)：%w", string(bs), err)
	}
	rsp = &OrderResponse{
		ResponseCode: m.Get("responseCode"),
		PayUrl:       m.Get("tn"),
	}
	return rsp, nil
}


func (x Client) doOrder(bm xthird.BodyMap) (bs []byte, err error) {

	payChannelType := "13"
	appId := x.WechatAppId
	key := x.WechatAppKey

	outputType := "2"
	//userIp := "101.80.79.107"

	bm.Set("appId", appId)
	bm.Set("deviceType", "0601")
	//bm.Set("frontNotifyUrl", "https://www.baidu.com/")
	bm.Set("funcode", "WP001")
	bm.Set("mhtCharset", "UTF-8")
	bm.Set("mhtCurrencyType", "156")
	//bm.Set("mhtOrderAmt", "1") // 金额
	//bm.Set("mhtOrderDetail", "订单详情")
	//bm.Set("mhtOrderName", "mhtOrderName")
	//bm.Set("mhtOrderNo", "1111144441") // 订单编号
	bm.Set("mhtOrderStartTime", time.Now().Format("20060102150405"))
	bm.Set("mhtOrderTimeOut", "3600")
	bm.Set("mhtOrderType", "01")
	bm.Set("mhtReserved", "test")
	//bm.Set("notifyUrl", "https://miniprog.pkey.cn/ym/api/v1/open/notifyPay/ttttttt")
	bm.Set("outputType", outputType) // 0 直接跳转微信支付页面，已进行页面封装; 2 返回 weixin://支付链接，需商户在前端使用 html 中的 a 标签调起支付;
	bm.Set("payChannelType", payChannelType)
	bm.Set("version", "1.0.0")
	//bm.Set("consumerCreateIp", userIp)
	bm.Set("mhtSignType", "MD5")

	signature := Sign(bm, key)
	bm.Set("mhtSignature", signature)
	param := bm.FormatURLParam()
	headerLocation := TradeUrl + param

	httpClient := xhttp.NewClient()
	httpClient.Header.Add("location", headerLocation)

	res, bs, errs := httpClient.Type(xhttp.TypeFormData).Post(TradeUrl).SendString(param).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil

}
