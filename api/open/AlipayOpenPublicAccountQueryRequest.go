package open

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 查询绑定商户会员号
// 当用户成为商户的关注用户之后，则商户可以通过本接口查询关注者的绑定账户，以便补全异常情况下的单边账户数据。
type AlipayOpenPublicAccountQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                        `json:"terminal_type"`
	TerminalInfo string                                        `json:"terminal_info"`
	ProdCode     string                                        `json:"prod_code"`
	NotifyURL    string                                        `json:"notify_url"`
	ReturnURL    string                                        `json:"return_url"`
	BizContent   AlipayOpenPublicAccountQueryRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicAccountQueryRequestBizContent struct {
	UserId string `json:"user_id"` // 支付宝账号userid，2088开头长度为16位的字符串
}

func (this *AlipayOpenPublicAccountQueryRequest) GetAPImethodName() string {
	return "alipay.open.public.account.query"
}

func (this *AlipayOpenPublicAccountQueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenPublicAccountQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenPublicAccountQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenPublicAccountQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenPublicAccountQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenPublicAccountQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenPublicAccountQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenPublicAccountQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
