package yyb

import (
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
	"reflect"
	"testing"
	"time"
)

func TestClient_doAuthToken(t *testing.T) {
	type fields struct {
		WechatAppId     string
		QQAppId         string
		SanBoxAppKey    string
		OnlineAppKey    string
		WechatAppSecret string
	}
	type args struct {
		bm        xthird.BodyMap
		loginType string
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
				WechatAppId:     "WechatAppId",
				QQAppId:         "QQAppId",
				SanBoxAppKey:    "SanBoxAppKey",
				OnlineAppKey:    "OnlineAppKey",
				WechatAppSecret: "WechatAppSecret",
			},
			args: args{
				bm: map[string]interface{}{
					"openid": "AEB75BAFCCC167BB2D9491E5325C8B06",
					"openkey": "2FFC7BCC553B94801B2C8D62E9E9778D",
					"timestamp": gofunc.Int64ToString(time.Now().Unix()),
				},
				loginType: QQ,
			},
		},
		{
			name: "testCase-02",
			fields: fields{
				WechatAppId:     "WechatAppId",
				QQAppId:         "QQAppId",
				SanBoxAppKey:    "SanBoxAppKey",
				OnlineAppKey:    "OnlineAppKey",
				WechatAppSecret: "WechatAppSecret",
			},
			args: args{
				bm: map[string]interface{}{
					"openid": "oogo75lX1Ce9tMuU_iq4qJsTRhkQ",
					"openkey": "45_8ADPFYpabLEZJYn7j_CjclCs48eaCAA9vfAvyrCo8-ROABMTZL0ZNdojIWGG-HJPHPb-PpMY0J5TjX9hUovRgdkzkIgik3127aiVjPvMmnc",
					"timestamp": gofunc.Int64ToString(time.Now().Unix()),
				},
				loginType: WECHAT,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				WechatAppId:     tt.fields.WechatAppId,
				QQAppId:         tt.fields.QQAppId,
				SanBoxAppKey:    tt.fields.SanBoxAppKey,
				OnlineAppKey:    tt.fields.OnlineAppKey,
				WechatAppSecret: tt.fields.WechatAppSecret,
			}
			gotBs, err := c.doAuthToken(tt.args.bm, tt.args.loginType)
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
		WechatAppId     string
		QQAppId         string
		SanBoxAppKey    string
		OnlineAppKey    string
		WechatAppSecret string
	}
	type args struct {
		bm        xthird.BodyMap
		loginType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRsp *TokenAuthResponse
		wantErr bool
	}{
		{
			name: "testCase-01",
			fields: fields{
				WechatAppId:     "WechatAppId",
				QQAppId:         "QQAppId",
				SanBoxAppKey:    "SanBoxAppKey",
				OnlineAppKey:    "OnlineAppKey",
				WechatAppSecret: "WechatAppSecret",
			},
			args: args{
				bm: map[string]interface{}{
					"openid": "AEB75BAFCCC167BB2D9491E5325C8B06",
					"openkey": "2FFC7BCC553B94801B2C8D62E9E9778D",
					"timestamp": gofunc.Int64ToString(time.Now().Unix()),
				},
				loginType: QQ,
			},

		},
		{
			name: "testCase-02",
			fields: fields{
				WechatAppId:     "WechatAppId",
				QQAppId:         "QQAppId",
				SanBoxAppKey:    "SanBoxAppKey",
				OnlineAppKey:    "OnlineAppKey",
				WechatAppSecret: "WechatAppSecret",
			},
			args: args{
				bm: map[string]interface{}{
					"openid": "oogo75lX1Ce9tMuU_iq4qJsTRhkQ",
					"openkey": "45_8ADPFYpabLEZJYn7j_CjclCs48eaCAA9vfAvyrCo8-ROABMTZL0ZNdojIWGG-HJPHPb-PpMY0J5TjX9hUovRgdkzkIgik3127aiVjPvMmnc",
					"timestamp": gofunc.Int64ToString(time.Now().Unix()),
				},
				loginType: WECHAT,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				WechatAppId:     tt.fields.WechatAppId,
				QQAppId:         tt.fields.QQAppId,
				SanBoxAppKey:    tt.fields.SanBoxAppKey,
				OnlineAppKey:    tt.fields.OnlineAppKey,
				WechatAppSecret: tt.fields.WechatAppSecret,
			}
			gotRsp, err := c.AuthToken(tt.args.bm, tt.args.loginType)
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