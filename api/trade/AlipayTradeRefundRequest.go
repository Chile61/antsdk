package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradeRefundRequest alipay.trade.refund(统一收单交易退款接口)
// 当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，支付宝将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。
// 交易超过约定时间（签约时设置的可退款时间）的订单无法进行退款 支付宝退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。
// 一笔退款失败后重新提交，要采用原来的退款单号。总退款金额不能超过用户实际支付金额
type AlipayTradeRefundRequest struct {
	api.PublicAppAuthToken
	OutTradeNO   string `json:"out_trade_no"`   // 特殊可选	订单支付时传入的商户订单号,不能和 trade_no同时为空。
	TradeNO      string `json:"trade_no"`       // 特殊可选 支付宝交易号，和商户订单号不能同时为空
	RefundAmount string `json:"refund_amount"`  // 必选	需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundReason string `json:"refund_reason"`  // 可选 退款的原因说明
	OutRequestNO string `json:"out_request_no"` // 可选 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	OperatorID   string `json:"operator_id"`    // 可选 商户的操作员编号
	StoreID      string `json:"store_id"`       // 可选 商户的门店编号
	TerminalID   string `json:"terminal_id"`    // 可选 商户的终端编号

}

// GetAPImethodName GetAPImethodName
func (requ *AlipayTradeRefundRequest) GetAPImethodName() string {
	return "alipay.trade.refund"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayTradeRefundRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayTradeRefundRequest) IsNeedBizContent() bool {
	return true
}