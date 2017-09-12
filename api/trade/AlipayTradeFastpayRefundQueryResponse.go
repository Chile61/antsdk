package trade

import "../../api"

// AlipayTradeFastpayRefundQueryResponse AlipayTradeFastpayRefundQueryResponse
type AlipayTradeFastpayRefundQueryResponse struct {
	E            api.Exception
	TradeNO      string `json:"trade_no"`       // 支付宝交易号
	OutTradeNO   string `json:"out_trade_no"`   // 创建交易传入的商户订单号
	OutRequestNO string `json:"out_request_no"` // 本笔退款对应的退款请求号
	RefundReason string `json:"refund_reason"`  // 发起退款时，传入的退款原因
	TotalAmount  string `json:"total_amount"`   // 该笔退款所对应的交易的订单金额
	RefundAmount string `json:"refund_amount"`  // 本次退款请求，对应的退款金额

}

// SetTags SetTags
func (resp *AlipayTradeFastpayRefundQueryResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_fastpay_refund_query_response"
	exceptionTag = "alipay_trade_fastpay_refund_query_response"
	e = &resp.E
	return
}
