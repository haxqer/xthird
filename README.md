# xthird

对接第三方渠道的登陆+支付，对接方式以 `联运游戏` 为主

---
# 一、安装

```bash
$ go get -u github.com/haxqer/xthird
```

---
## todo (附上渠道文档地址)

- [x] [OPPO](https://open.oppomobile.com/)
- [ ] 微信 
- [ ] 支付宝
- [ ] 苹果
- [ ] VIVO
- [ ] 华为
- [ ] 小米
- [ ] 魅族
- [ ] 应用宝

---
# 二、使用说明

## OPPO 示例
OPPO 登陆:
```go
import (
	"github.com/haxqer/xthird/oppo"
)

func main()  {

    client := oppo.NewClient("YourAppId", "YourAppKey", "YourAppSec")
    
    body := make(xthird.BodyMap)
    body.Set("oss_id", "oss_id - Get from client of sdk")
    body.Set("token", "token - Get from client of sdk")
    
    oppoRsp, err := client.Login(body)
    if err != nil {
        // todo 处理错误逻辑
        // 如返回客户端 `参数错误` 或者 `登陆失败`
        return
    }
    
    // 需要判断 oppoRsp.ResultCode 是否等于 200
    if oppoRsp.ResultCode != "200" {
        // todo 处理错误逻辑
        // 如返回客户端 `登陆失败`
        return
    }
    
    fmt.Printf("result_code:%s\n", oppoRsp.ResultCode)
    fmt.Printf("result_msg:%s\n", oppoRsp.ResultMsg)
    fmt.Printf("user_name:%s\n", oppoRsp.UserName)
    fmt.Printf("email:%s\n", oppoRsp.Email)
    fmt.Printf("phone:%s\n", oppoRsp.MobileNumber)
    fmt.Printf("sso_id:%s\n", oppoRsp.SsoId)
}


```



OPPO 登陆 输出 demo:
```bash
result_code:200
result_msg:正常
user_name:用户 01234566
email:
phone:
sso_id:1234
```



OPPO 回调

```go
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/haxqer/xthird/oppo"
)

func notify(c *gin.Context)  {
	// OPPO公钥. 在官方给的 demo 中. 无需修改
	oppoPublicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCmreYIkPwVovKR8rLHWlFVw7YDfm9uQOJKL89Smt6ypXGVdrAKKl0wNYc3/jecAoPi2ylChfa2iRu5gunJyNmpWZzlCNRIau55fxGW0XEu553IiprOZcaw5OuYGlf60ga8QT6qToP0/dpiL/ZbmNUO9kUhosIjEu22uFgR+5cYyQIDAQAB"

	// 解析请求参数
	bodyMap, err := oppo.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		// 解析失败, 处理错误逻辑
		return
	}

	// 验签
	err = oppo.VerifySign(oppoPublicKey, bodyMap)
	if err != nil {
		// 验签失败, 处理错误逻辑
		return
	}


	fmt.Printf("oppo order id:%s\n", bodyMap.GetString("notifyId"))
	fmt.Printf("game order id:%s\n", bodyMap.GetString("partnerOrder")) // 游戏订单id。 如果存在聚合平台，则为 聚合平台订单id
	// 需要严格校验 价格/金额
	fmt.Printf("price:%s\n", bodyMap.GetString("price")) // 商品价格(以分为单位)
	fmt.Printf("product name:%s\n", bodyMap.GetString("productName")) // 商品名称(客户端上传)
	fmt.Printf("product desc:%s\n", bodyMap.GetString("productDesc")) // 商品名称(客户端上传)
	fmt.Printf("count:%s\n", bodyMap.GetString("count")) // 商品数量(一般为 1)
}
```

oppo 回调 form 示例:
```bash
{
    "adId": "",
    "attach": "自定义字段",
    "channel": "1",
    "count": "1",
    "notifyId": "GC2021060901398888888",
    "partnerOrder": "11111111",
    "price": "1",
    "productDesc": "55555",
    "productName": "300钻石",
    "sign": "XXXX/b/Uhmzyvpj3Wta4BQdyvv8wJKht1q6FfelcIzuU6SU7sTfAF1J0hkc=",
    "userId": "111111"
}
```