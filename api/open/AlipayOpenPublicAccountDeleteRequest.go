package open

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 解除绑定商户会员号
// 在支付宝服务窗中目前一共有三种解除绑定商户会员号的场景，具体如下： 用户取消关注服务窗 用户在服务窗内手动执行解除绑定操作 开发者调用解除绑定商户会员号接口解除绑定
type AlipayOpenPublicAccountDeleteRequest struct {
	api.IAlipayRequest
	TerminalType string                                         `json:"terminal_type"`
	TerminalInfo string                                         `json:"terminal_info"`
	ProdCode     string                                         `json:"prod_code"`
	NotifyURL    string                                         `json:"notify_url"`
	ReturnURL    string                                         `json:"return_url"`
	BizContent   AlipayOpenPublicAccountDeleteRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicAccountDeleteRequestBizContent struct {
	AgreementId   string `json:"agreement_id"`    // 协议号，商户会员在支付宝服务窗账号中的唯一标识，与bind_account_no不能同时为空
	BindAccountNo string `json:"bind_account_no"` // 绑定帐号，建议在开发者的系统中保持唯一性，与agreement_id不能同时为空
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetAPImethodName() string {
	return "alipay.open.public.account.delete"
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenPublicAccountDeleteRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenPublicAccountDeleteRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
