package vivo

import (
	"encoding/json"
	"github.com/haxqer/xthird"
	"reflect"
	"testing"
)

func TestAuthToken(t *testing.T) {
	type args struct {
		bm xthird.BodyMap
	}
	tests := []struct {
		name    string
		args    args
		wantRsp *TokenAuthResponse
		wantErr bool
	}{
		{
			name: "testCase-01",
			args: args{
				bm: map[string]interface{}{
					"opentoken" : "_STV1_797e3324f7e3f1a3_797e3324f7e3f1a3_8db97942_Awykia3hpb90kcu3l",
				},
			},
			wantRsp: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := AuthToken(tt.args.bm)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				if gotRsp != nil  {
					marshal, _ := json.Marshal(gotRsp)
					println(string(marshal))
				}
				t.Errorf("AuthToken() gotRsp = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}

func TestFormatURLParam(t *testing.T) {
	type args struct {
		body xthird.BodyMap
	}
	tests := []struct {
		name         string
		args         args
		wantUrlParam string
	}{
		{name: "testCase-01", args: args{body: map[string]interface{}{
			"opentoken" : "_STV1_797e3324f7e3f1a3_797e3324f7e3f1a3_8db97942_Abbccayhpb90kvd3m",
			"123" : "123",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUrlParam := FormatURLParam(tt.args.body); gotUrlParam != tt.wantUrlParam {
				t.Errorf("FormatURLParam() = %v, want %v", gotUrlParam, tt.wantUrlParam)
			}
		})
	}
}