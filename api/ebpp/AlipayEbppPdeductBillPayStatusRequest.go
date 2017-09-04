package ebpp

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 缴费直连代扣订单支付状态查询
type AlipayEbppPdeductBillPayStatusRequest struct {
	api.IAlipayRequest
	TerminalType string `json:"terminal_type"`
	TerminalInfo string `json:"terminal_info"`
	ProdCode     string `json:"prod_code"`
	NotifyURL    string `json:"notify_url"`
	ReturnURL    string `json:"return_url"`
	OutOrderNo   string `json:"out_order_no"` // 商户代扣业务流水
	AgreementId  string `json:"agreement_id"` // 支付宝用户ID
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetApiMethodName() string {
	return "alipay.ebpp.pdeduct.bill.pay.status"
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayEbppPdeductBillPayStatusRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayEbppPdeductBillPayStatusRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("out_order_no", this.OutOrderNo)
	txtParams.Put("agreement_id", this.AgreementId)
	return txtParams
}
