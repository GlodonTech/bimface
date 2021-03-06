package http

import (
	"encoding/base64"

	"github.com/imroc/req"
)

//Headers ***
type Headers struct {
	Header req.Header // = make(Header)
}

//NewHeaders ***
func NewHeaders() *Headers {
	o := &Headers{
		Header: make(req.Header),
	}
	o.Header["Content-Type"] = "application/json"

	return o
}

//AddBasicAuthHeader ***
func (o *Headers) AddBasicAuthHeader(appKey string, appSecret string) {
	key := appKey + ":" + appSecret
	authorization := "Basic" + " " + base64.StdEncoding.EncodeToString([]byte(key))
	o.Header["Authorization"] = authorization
}

//AddOAuth2Header ***
func (o *Headers) AddOAuth2Header(token string) {
	o.Header["Authorization"] = "Bearer " + token
}
