package oppo

import (
	"github.com/haxqer/xthird"
	"testing"
)

func TestVerifySign(t *testing.T) {
	type args struct {
		oppoPayPublicKey string
		bm               xthird.BodyMap
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "testCase-01",
			args: args{
				oppoPayPublicKey: "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCmreYIkPwVovKR8rLHWlFVw7YDfm9uQOJKL89Smt6ypXGVdrAKKl0wNYc3/jecAoPi2ylChfa2iRu5gunJyNmpWZzlCNRIau55fxGW0XEu553IiprOZcaw5OuYGlf60ga8QT6qToP0/dpiL/ZbmNUO9kUhosIjEu22uFgR+5cYyQIDAQAB",
				bm:               xthird.BodyMap{
					"adId": "",
					"attach": "自定义字段",
					"channel": "1",
					"count": "1",
					"notifyId": "notifyId",
					"partnerOrder": "partnerOrder",
					"price": "1",
					"productDesc": "55555",
					"productName": "300钻石",
					"sign": "sign",
					"userId": "userId",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifySign(tt.args.oppoPayPublicKey, tt.args.bm); (err != nil) != tt.wantErr {
				t.Errorf("VerifySign() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}