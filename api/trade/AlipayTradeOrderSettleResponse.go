package trade

import "../../api"

// AlipayTradeOrderSettleResponse AlipayTradeOrderSettleResponse
type AlipayTradeOrderSettleResponse struct {
	E       api.Exception
	TradeNO string `json:"trade_no"` // 支付宝交易号

}

// SetTags SetTags
func (resp *AlipayTradeOrderSettleResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_order_settle_response"
	exceptionTag = "alipay_trade_order_settle_response"
	e = &resp.E
	return
}
