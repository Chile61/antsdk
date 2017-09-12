package trade

import "../../api"

// AlipayTradeCloseResponse AlipayTradeCloseResponse
type AlipayTradeCloseResponse struct {
	E          api.Exception
	TradeNO    string `json:"trade_no"`     // 支付宝交易号
	OutTradeNO string `json:"out_trade_no"` // 创建交易传入的商户订单号

}

// SetTags SetTags
func (resp *AlipayTradeCloseResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_close_response"
	exceptionTag = "alipay_trade_close_response"
	e = &resp.E
	return
}
