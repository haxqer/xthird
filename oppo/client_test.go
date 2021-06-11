package oppo

import (
	"github.com/haxqer/xthird"
	"reflect"
	"testing"
)

func TestClient_Login(t *testing.T) {
	type fields struct {
		AppId     string
		AppKey    string
		AppSecret string
	}
	type args struct {
		bm xthird.BodyMap
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantOppoRsp *LoginResponse
		wantErr     bool
	}{
		{
			name: "testCase-01",
			fields: fields{
				AppId:     "AppId",
				AppKey:    "AppKey",
				AppSecret: "AppSecret",
			},
			args: args{
				bm: xthird.BodyMap{
					"sso_id": "sso_id",
					"token": "token",
				},
			},
			wantOppoRsp: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Client{
				AppId:     tt.fields.AppId,
				AppKey:    tt.fields.AppKey,
				AppSecret: tt.fields.AppSecret,
			}
			gotOppoRsp, err := o.Login(tt.args.bm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOppoRsp, tt.wantOppoRsp) {
				t.Errorf("Login() gotOppoRsp = %v, want %v", gotOppoRsp, tt.wantOppoRsp)
			}
		})
	}
}
