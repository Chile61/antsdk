package trade

import "../../api"

// AlipayTradeCancelResponse AlipayTradeCancelResponse
type AlipayTradeCancelResponse struct {
	E          api.Exception
	TradeNO    string `json:"trade_no"`     // 支付宝交易号
	OutTradeNO string `json:"out_trade_no"` // 商户订单号
	RetryFlag  string `json:"retry_flag"`   // 是否需要重试 示例值N
	Action     string `json:"action"`       // 	本次撤销触发的交易动作 close：关闭交易，无退款 refund：产生了退款

}

// SetTags SetTags
func (resp *AlipayTradeCancelResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_cancel_response"
	exceptionTag = "alipay_trade_cancel_response"
	e = &resp.E
	return
}
