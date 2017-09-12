package util

import "../../api"

// AlipaySystemOauthTokenRequest alipay.system.oauth.token(换取授权访问令牌)
// 换取授权访问令牌
type AlipaySystemOauthTokenRequest struct {
	api.PublicAppAuthToken
	GrantType    string `json:"grant_type"`    // 必选 值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code"`          // 可选 授权码，用户对应用授权后得到。
	RefreshToken string `json:"refresh_token"` // 可选 刷新令牌，上次换取访问令牌时得到。见出参的refresh_token字段
}

// GetAPImethodName GetAPImethodName
func (requ *AlipaySystemOauthTokenRequest) GetAPImethodName() string {
	return "alipay.system.oauth.token"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipaySystemOauthTokenRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipaySystemOauthTokenRequest) IsNeedBizContent() bool {
	return false
}
