package ebpp

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 缴费直连代扣取消签约
// 缴费直连代扣，用户取消签约接口
type AlipayEbppPdeductSignCancelRequest struct {
	api.IAlipayRequest
	TerminalType     string `json:"terminal_type"`
	TerminalInfo     string `json:"terminal_info"`
	ProdCode         string `json:"prod_code"`
	NotifyURL        string `json:"notify_url"`
	ReturnURL        string `json:"return_url"`
	UserId           string `json:"user_id"`
	AgreementId      string `json:"agreement_id"`
	AgentChannel     string `json:"agent_channel"`
	AgentCode        string `json:"agent_code"`
	PayPasswordToken string `json:"pay_password_token"`
}

func (this *AlipayEbppPdeductSignCancelRequest) GetAPImethodName() string {
	return "alipay.ebpp.pdeduct.sign.cancel"
}

func (this *AlipayEbppPdeductSignCancelRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayEbppPdeductSignCancelRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayEbppPdeductSignCancelRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayEbppPdeductSignCancelRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayEbppPdeductSignCancelRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayEbppPdeductSignCancelRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayEbppPdeductSignCancelRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayEbppPdeductSignCancelRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("user_id", this.UserId)
	txtParams.Put("agreement_id", this.AgreementId)
	txtParams.Put("agent_channel", this.AgentChannel)
	txtParams.Put("agent_code", this.AgentCode)
	txtParams.Put("pay_password_token", this.PayPasswordToken)
	return txtParams
}
