package xiaomi

import (
	"github.com/haxqer/xthird"
	"testing"
)

func TestSign(t *testing.T) {
	type args struct {
		bm  xthird.BodyMap
		key string
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
					"id":      "1234",
					"orderId": "12312312",
					"status":  1,
				},
				key: "2CmeD1ABmr5u6/1VEHjY7g==",
			},
			want: "6782e2c31cec7ef777c04f3371abf4175ddbd469",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.bm, tt.args.key); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
