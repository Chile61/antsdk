package trade

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 统一收单交易退款查询
// 商户可使用该接口查询自已通过alipay.trade.refund提交的退款请求是否执行成功。
type AlipayTradeFastpayRefundQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                         `json:"terminal_type"`
	TerminalInfo string                                         `json:"terminal_info"`
	ProdCode     string                                         `json:"prod_code"`
	NotifyURL    string                                         `json:"notify_url"`
	ReturnURL    string                                         `json:"return_url"`
	BizContent   AlipayTradeFastpayRefundQueryRequestBizContent `json:"biz_content"`
}

type AlipayTradeFastpayRefundQueryRequestBizContent struct {
	TradeNo      string `json:"trade_no"`       // 支付宝交易号，和商户订单号不能同时为空
	OutTradeNo   string `json:"out_trade_no"`   // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	OutRequestNo string `json:"out_request_no"` // 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetAPImethodName() string {
	return "alipay.trade.fastpay.refund.query"
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayTradeFastpayRefundQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayTradeFastpayRefundQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
