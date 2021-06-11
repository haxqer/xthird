package ymmonster

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
					"appid": "888888",
					"un": "develope",
					"pwd": "123456",
					"other": "JYKO66BNN",
				},
				key: "E10ADC3949BA59ABBE56E057F20F883E",
			},
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
