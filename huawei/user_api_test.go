package huawei

import (
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
			args: args{bm: map[string]interface{}{
				"access_token": "xxxxx/l6PNWz5vzv3FylBFT+w==",
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

func Test_doUserPost(t *testing.T) {
	type args struct {
		bm  xthird.BodyMap
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantBs  []byte
		wantErr bool
	}{
		{
			name: "testCase-01",
			args: args{
				bm: map[string]interface{}{
					"access_token": "xxxxxx/l6PNWz5vzv3FylBFT+w==",
					"open_id":      "OPENID",
				},
				url: AuthTokenUrl,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBs, err := doUserPost(tt.args.bm, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("doUserPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBs, tt.wantBs) {
				t.Errorf("doUserPost() gotBs = %s, want %v", gotBs, tt.wantBs)
			}
		})
	}

}