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
			want: "7728023d3273211147a3dfb8d2f0dd5947acfb8b",
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
					"appid": "888888",
					"un": "develope",
					"pwd": "123456",
					"other": "JYKO66BNN",
					"signature": "7728023d3273211147a3dfb8d2f0dd5947acfb8b",
				},
				key: "E10ADC3949BA59ABBE56E057F20F883E",
			},
			want: true,
		},
		{
			name: "testCase-02",
			args: args{
				bm: map[string]interface{}{
					"cp_ordernum": "1193844986517397504",
					"order_env": "0",
					"ext_info": "",
					"userid": "15621",
					"delive_time": "1623410076000",
					"delive_money": "1",
					"sdk_ordernum": "20210611191347676001617794016884",
					"signature": "383d4b0ea8e76c5f58d46165db06cfde520b0619",
				},
				//key: "99AEF2712A81005B2F4EC1B92BED545D",
				key: "37026B55007DCF0800D99EF68B65015A",
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