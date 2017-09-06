package util

import "github.com/vanishs/antsdk/api"

// AlipayOpenAuthTokenAppResponse AlipayOpenAuthTokenAppResponse
type AlipayOpenAuthTokenAppResponse struct {
	E               api.Exception
	UserID          string `json:"user_id"`           // 授权商户的user_id
	AuthAppID       string `json:"auth_app_id"`       // 授权商户的appid
	AppAuthToken    string `json:"app_auth_token"`    // 应用授权令牌
	AppRefreshToken string `json:"app_refresh_token"` // 刷新令牌
	ExpiresIn       int64  `json:"expires_in"`        // 应用授权令牌的有效时间（从接口调用时间作为起始时间），单位到秒
	ReExpiresIn     int64  `json:"re_expires_in"`     // 刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒
}

// SetTags SetTags
func (resp *AlipayOpenAuthTokenAppResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_open_auth_token_app_response"
	exceptionTag = "alipay_open_auth_token_app_response"
	e = &resp.E
	return
}