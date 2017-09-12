package util

import "github.com/vanishs/antsdk/api"

// AlipayUserInfoAuthResponse AlipayUserInfoAuthResponse
type AlipayUserInfoAuthResponse struct {
	E api.Exception
}

// SetTags SetTags
func (resp *AlipayUserInfoAuthResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_user_info_auth_response"
	exceptionTag = "alipay_user_info_auth_response"
	e = &resp.E
	return
}