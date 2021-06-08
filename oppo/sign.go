package oppo

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/haxqer/gofunc"
	"github.com/haxqer/xthird"
	"hash"
	"math/rand"
	"net/url"
	"time"
)

func GenLoginBaseStr(bm xthird.BodyMap, appKey, appSecret string) (string, string) {
	baseStr := fmt.Sprintf("oauthConsumerKey=%s&oauthToken=%s&oauthSignatureMethod=HMAC-SHA1&oauthTimestamp=%d&oauthNonce=%d&oauthVersion=1.0&",
		appKey, url.QueryEscape(bm["token"].(string)), time.Now().Unix(), rand.Int31n(100000000))

	var h hash.Hash
	h = hmac.New(sha1.New, []byte(appSecret + "&" ))
	h.Write([]byte(baseStr))

	sign := url.QueryEscape(gofunc.Base64Encode(string(h.Sum(nil))))
	return baseStr, sign
}

