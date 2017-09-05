package ebpp

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 查询账单
type AlipayEbppBillGetRequest struct {
	api.IAlipayRequest
	TerminalType    string `json:"terminal_type"`
	TerminalInfo    string `json:"terminal_info"`
	ProdCode        string `json:"prod_code"`
	NotifyURL       string `json:"notify_url"`
	ReturnURL       string `json:"return_url"`
	OrderType       string `json:"order_type"`
	MerchantOrderNo string `json:"merchant_order_no"`
}

func (this *AlipayEbppBillGetRequest) GetAPImethodName() string {
	return "alipay.ebpp.bill.get"
}

func (this *AlipayEbppBillGetRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayEbppBillGetRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayEbppBillGetRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayEbppBillGetRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayEbppBillGetRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayEbppBillGetRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayEbppBillGetRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayEbppBillGetRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("order_type", this.OrderType)
	txtParams.Put("merchant_order_no", this.MerchantOrderNo)
	return txtParams
}
