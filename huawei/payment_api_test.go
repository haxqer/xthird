package huawei

import (
	"github.com/haxqer/xthird"
	"reflect"
	"testing"
)

func TestClient_OrderVerify(t *testing.T) {
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
		wantRsp *OrderResponse
		wantErr bool
	}{
		{
			name: "testCase-01",
			fields: fields{
				AppId:     "AppId",
				AppKey:    "AppKey",
				AppSecret: "AppSecret",
			},
			args: args{map[string]interface{}{
				"access_token": "xxxxxx/PG5xy8iEocIIwfidW1dqK+gLuEsXgYJD8X0ckO/qzO+3WQeyFKgT7GtKsyRN78nf+hLDbEWpJAOvidZtiJKLOplmQ==",
				"purchaseToken": "xxxxxxx.1.100883123",
				"productId": "test001",
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
			gotRsp, err := c.OrderVerify(tt.args.bm)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderVerify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("OrderVerify() gotRsp = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}