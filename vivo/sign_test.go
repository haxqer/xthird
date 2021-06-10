package vivo

import (
	"github.com/haxqer/xthird"
	"testing"
)

func TestVerifySign(t *testing.T) {
	type args struct {
		bm  xthird.BodyMap
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testCase-01",
			args: args{
				bm: map[string]interface{}{
					"appId":"111",
					"cpId":"11",
					"cpOrderNumber":"111",
					"extInfo":"扩展参数",
					"orderAmount":"1",
					"orderNumber":"11",
					"payTime":"20210610213219",
					"respCode":"200",
					"respMsg":"交易成功",
					"signMethod":"MD5",
					"signature":"111",
					"tradeStatus":"0000",
					"tradeType":"01",
					"uid":"111",

				},
				key: "1111",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifySign(tt.args.bm, tt.args.key); got != tt.want {
				t.Errorf("VerifySign() = %v, want %v", got, tt.want)
			}
		})
	}
}
