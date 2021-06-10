package huawei

import (
	"github.com/haxqer/xthird"
	"reflect"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	type args struct {
		bm xthird.BodyMap
	}
	tests := []struct {
		name    string
		args    args
		wantRsp *GetUserInfoResponse
		wantErr bool
	}{
		{
			name: "testCase-01",
			args: args{bm: map[string]interface{}{
				"access_token": "",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := GetUserInfo(tt.args.bm)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("GetUserInfo() gotRsp = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}

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
			args: args{bm: map[string]interface{}{
				"access_token": "",
			}},
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
				t.Errorf("AuthToken() gotRsp = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
