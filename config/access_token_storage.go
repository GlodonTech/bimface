package config

import (
	"bimface/bean/response"
	"time"
)

//AccessTokenStorage ***
type AccessTokenStorage struct {
	accessToken *response.AccessToken
}

//NewAccessTokenStorage ***
func NewAccessTokenStorage() *AccessTokenStorage {
	o := &AccessTokenStorage{
		accessToken: nil,
	}
	return o
}

//Put ***
func (o *AccessTokenStorage) Put(accessToken *response.AccessToken) {
	o.accessToken = accessToken
}

//Get ***
func (o *AccessTokenStorage) Get() *response.AccessToken {

	if o.accessToken == nil {
		return nil
	}

	if maybeExpire(o.accessToken.ExpireTime) {
		return nil
	}

	return o.accessToken
}

//MaybeExpire ***
func maybeExpire(expireTime string) bool {
	expire, err := time.ParseInLocation("2006-01-02 15:04:05", expireTime, time.Local)
	if err == nil {
		return expire.Before(time.Now())
	}

	return true
}
