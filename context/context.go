package context

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Context struct {
	Timestamp     int64  //当前时间戳
	Envid         string //环境id
	SecretId      string //id
	SecretKey     string //secret
	Authorization string
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])

}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))

}

//GetAuth 获取凭证
func (ctx *Context) GetAuth() error {
	//host := "cvm.tencentcloudapi.com"
	algorithm := "TC3-HMAC-SHA256"
	service := "tcb"
	version := "1.0"

	//var timestamp int64 = time.Now().Unix()
	var timestamp int64 = ctx.Timestamp
	// step 1: build canonical request string
	httpRequestMethod := "POST"
	canonicalURI := "//api.tcloudbase.com/"
	canonicalQueryString := ""
	canonicalHeaders := "content-type:application/json; charset=utf-8\nhost:api.tcloudbase.com\n"
	signedHeaders := "content-type;host"

	hashedRequestPayload := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)

	// step 2: build string to sign
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")

	//date := "2020-09-16"

	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest)

	// step 3: sign string
	secretDate := hmacsha256(date, "TC3"+ctx.SecretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))

	// step 4: build authorization
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		ctx.SecretId,
		credentialScope,
		signedHeaders,
		signature)
	ctx.Authorization = version + " " + authorization

	return nil
}
