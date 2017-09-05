package market

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 服务市场商户确认订购通知
// 服务市场当商户选择服务商提供产品并订购确认时，通知服务商订单消息。服务商可以通过通知的消息内容回查该订单明细。回查接口（alipay.open.servicemarket.order.query）
type AlipayOpenServicemarketOrderNotifyRequest struct {
	api.IAlipayRequest
	TerminalType string `json:"terminal_type"`
	TerminalInfo string `json:"terminal_info"`
	ProdCode     string `json:"prod_code"`
	NotifyURL    string `json:"notify_url"`
	ReturnURL    string `json:"return_url"`
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetAPImethodName() string {
	return "alipay.open.servicemarket.order.notify"
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenServicemarketOrderNotifyRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	return txtParams
}
