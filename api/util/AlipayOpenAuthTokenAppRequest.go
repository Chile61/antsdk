package util

// AlipayOpenAuthTokenAppRequest alipay.open.auth.token.app(换取应用授权令牌)
// 换取应用授权令牌。在应用授权的场景下，商户把名下应用授权给ISV后，支付宝会给ISV颁发应用授权码app_auth_code，ISV可通过获取到的app_auth_code换取app_auth_token。app_auth_code作为换取app_auth_token的票据，每次用户授权带上的app_auth_code将不一样，app_auth_code只能使用一次，一天（从当前时间算起的24小时）未被使用自动过期。 刷新应用授权令牌，ISV可通过获取到的refresh_token刷新app_auth_token，刷新后老的refresh_token会在一段时间后失效（失效时间为接口返回的re_expires_in）。
type AlipayOpenAuthTokenAppRequest struct {
	GrantType    string `json:"grant_type"`    // 必选 authorization_code表示换取app_auth_token。 refresh_token表示刷新app_auth_token。
	Code         string `json:"code"`          // 可选 授权码，如果grant_type的值为authorization_code。该值必须填写
	RefreshToken string `json:"refresh_token"` // 可选 刷新令牌，如果grant_type值为refresh_token。该值不能为空。该值来源于此接口的返回值app_refresh_token（至少需要通过grant_type=authorization_code调用此接口一次才能获取）
}

// GetAPImethodName GetAPImethodName
func (requ *AlipayOpenAuthTokenAppRequest) GetAPImethodName() string {
	return "alipay.open.auth.token.app"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayOpenAuthTokenAppRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayOpenAuthTokenAppRequest) IsNeedBizContent() bool {
	return true
}
