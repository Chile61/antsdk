package user

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// AlipayUserUserinfoShareRequest alipay.user.userinfo.share(支付宝钱包用户信息共享)
// 外部应用上架到支付宝钱包，当支付宝用户从钱包访问外部应用时，会跳转到外部应用并带上用户的授权码。 外部应用用授权码调用授权令牌交换API（alipay.system.oauth.token）可得到授权令牌。 用授权令牌调用此接口得到支付宝会员相关信息。 特别说明：此接口的不需要授权是指不需外部应用主动引导用户授权，支付宝钱包会在引导用户授权后， 带上授权码再跳转到外部应用。
type AlipayUserUserinfoShareRequest struct {
	api.IAlipayRequest
	TerminalType string `json:"terminal_type"`
	TerminalInfo string `json:"terminal_info"`
	ProdCode     string `json:"prod_code"`
	NotifyURL    string `json:"notify_url"`
	ReturnURL    string `json:"return_url"`
}

// GetApiMethodName GetApiMethodName
func (requ *AlipayUserUserinfoShareRequest) GetApiMethodName() string {
	return "alipay.user.userinfo.share"
}

// GetApiVersion GetApiVersion
func (requ *AlipayUserUserinfoShareRequest) GetApiVersion() string {
	return "1.0"
}

// GetTerminalType GetTerminalType
func (requ *AlipayUserUserinfoShareRequest) GetTerminalType() string {
	return requ.TerminalType
}

// GetTerminalInfo GetTerminalInfo
func (requ *AlipayUserUserinfoShareRequest) GetTerminalInfo() string {
	return requ.TerminalInfo
}

// GetNotifyURL GetNotifyURL
func (requ *AlipayUserUserinfoShareRequest) GetNotifyURL() string {
	return requ.NotifyURL
}

// GetReturnURL GetReturnURL
func (requ *AlipayUserUserinfoShareRequest) GetReturnURL() string {
	return requ.ReturnURL
}

// GetProdCode GetProdCode
func (requ *AlipayUserUserinfoShareRequest) GetProdCode() string {
	return requ.ProdCode
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayUserUserinfoShareRequest) IsNeedEncrypt() bool {
	return false
}

// GetTextParams GetTextParams
func (requ *AlipayUserUserinfoShareRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	return txtParams
}
