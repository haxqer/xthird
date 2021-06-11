package ymmonster

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
				AppId:     "20000013",
				AppKey:    "E297993F2BD63F642A9EAE1236856963",
				AppSecret: "15CA4512C592CBBAB94FBCEDA2A747EB",
			},
			args: args{map[string]interface{}{
				"userId": "15621",
				"token": "f1025f2695494191a0cef90739188970!20210611181615",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Client{
				AppId:     tt.fields.AppId,
				AppKey:    tt.fields.AppKey,
				AppSecret: tt.fields.AppSecret,
			}
			gotBs, err := x.doAuthToken(tt.args.bm)
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
				AppId:     "20000013",
				AppKey:    "E297993F2BD63F642A9EAE1236856963",
				AppSecret: "15CA4512C592CBBAB94FBCEDA2A747EB",
			},
			args: args{map[string]interface{}{
				"userId": "15621",
				"token": "f1025f2695494191a0cef90739188970!20210611181615",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Client{
				AppId:     tt.fields.AppId,
				AppKey:    tt.fields.AppKey,
				AppSecret: tt.fields.AppSecret,
			}
			gotRsp, err := x.AuthToken(tt.args.bm)
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

func TestClient_doAuthToken1(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Client{
				AppId:     tt.fields.AppId,
				AppKey:    tt.fields.AppKey,
				AppSecret: tt.fields.AppSecret,
			}
			gotBs, err := x.doAuthToken(tt.args.bm)
			if (err != nil) != tt.wantErr {
				t.Errorf("doAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBs, tt.wantBs) {
				t.Errorf("doAuthToken() gotBs = %v, want %v", gotBs, tt.wantBs)
			}
		})
	}
}