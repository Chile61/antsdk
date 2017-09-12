package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradeFastpayRefundQueryRequest alipay.trade.fastpay.refund.query(统一收单交易退款查询)
// 商户可使用该接口查询自已通过alipay.trade.refund提交的退款请求是否执行成功。
// 该接口的返回码10000，仅代表本次查询操作成功，不代表退款成功。
// 如果该接口返回了查询数据，则代表退款成功，如果没有查询到则代表未退款成功，可以调用退款接口进行重试。重试时请务必保证退款请求号一致。
type AlipayTradeFastpayRefundQueryRequest struct {
	api.PublicAppAuthToken
	TradeNO      string `json:"trade_no"`       // 特殊可选 支付宝交易号，和商户订单号不能同时为空
	OutTradeNO   string `json:"out_trade_no"`   // 特殊可选	订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	OutRequestNO string `json:"out_request_no"` // 必选	  请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
}

// GetAPImethodName GetAPImethodName
func (requ *AlipayTradeFastpayRefundQueryRequest) GetAPImethodName() string {
	return "alipay.trade.fastpay.refund.query"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayTradeFastpayRefundQueryRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayTradeFastpayRefundQueryRequest) IsNeedBizContent() bool {
	return true
}