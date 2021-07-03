package meizu

import (
	"github.com/haxqer/xthird"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	type args struct {
		bm        xthird.BodyMap
		appSecret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testCase-01",
			args: args{
				bm: map[string]interface{}{
					"app_id": "3369924",
					"cp_order_id": "1623824834",
					"uid": "171203910",
					"product_id": "0",
					"product_subject": "购买 N 枚金币",
					"product_body": "",
					"product_unit": "",
					"buy_amount": "1",
					"product_per_price": "0.01",
					"total_price": "0.01",
					"create_time": "1623824834",
					"pay_type": "0",
					"user_info": "",
					"sign_type": "md5",
				},
				appSecret: "xxxxxxxx",
			},
		},
	}

	println(time.Now().Unix())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.bm, tt.args.appSecret); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
