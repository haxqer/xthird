package meizu

import (
	"github.com/haxqer/xthird"
	"reflect"
	"testing"
)

func TestClient_doAuthToken(t *testing.T) {
	type fields struct {
		AppId     string
		AppKey    string
		AppSecret string
	}
	type args struct {
		bm xthird.BodyMap
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBs  []byte
		wantErr bool
	}{
		{
			name: "testCase-01",
			fields: fields{
				AppId:     "AppId",
				AppKey:    "AppKey",
				AppSecret: "AppSecret",
			},
			args: args{bm: map[string]interface{}{
				"uid": "xxxxxx",
				"session_id": "session_id",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AppId:     tt.fields.AppId,
				AppKey:    tt.fields.AppKey,
				AppSecret: tt.fields.AppSecret,
			}
			gotBs, err := c.doAuthToken(tt.args.bm)
			if (err != nil) != tt.wantErr {
				t.Errorf("doAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBs, tt.wantBs) {
				t.Errorf("doAuthToken() gotBs = %s, want %v", gotBs, tt.wantBs)
			}
		})
	}
}

func TestClient_AuthToken(t *testing.T) {
	type fields struct {
		AppId     string
		AppKey    string
		AppSecret string
	}
	type args struct {
		bm xthird.BodyMap
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRsp *AuthResponse
		wantErr bool
	}{
		{
			name: "testCase-01",
			fields: fields{
				AppId:     "AppId",
				AppKey:    "AppKey",
				AppSecret: "AppSecret",
			},
			args: args{bm: map[string]interface{}{
				"uid": "uid",
				"session_id": "session_id",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AppId:     tt.fields.AppId,
				AppKey:    tt.fields.AppKey,
				AppSecret: tt.fields.AppSecret,
			}
			gotRsp, err := c.AuthToken(tt.args.bm)
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