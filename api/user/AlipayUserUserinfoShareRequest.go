package user

import "github.com/vanishs/antsdk/api"

// AlipayUserUserinfoShareRequest alipay.user.userinfo.share(支付宝钱包用户信息共享)
// 外部应用上架到支付宝钱包，当支付宝用户从钱包访问外部应用时，会跳转到外部应用并带上用户的授权码。 外部应用用授权码调用授权令牌交换API（alipay.system.oauth.token）可得到授权令牌。 用授权令牌调用此接口得到支付宝会员相关信息。 特别说明：此接口的不需要授权是指不需外部应用主动引导用户授权，支付宝钱包会在引导用户授权后， 带上授权码再跳转到外部应用。
type AlipayUserUserinfoShareRequest struct {
	api.PublicAppAuthToken
	api.PublicAuthToken
}

// GetAPImethodName GetAPImethodName
func (requ *AlipayUserUserinfoShareRequest) GetAPImethodName() string {
	return "alipay.user.userinfo.share"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayUserUserinfoShareRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayUserUserinfoShareRequest) IsNeedBizContent() bool {
	return false
}