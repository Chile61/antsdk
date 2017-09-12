package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradeCreateResponse AlipayTradeCreateResponse
type AlipayTradeCreateResponse struct {
	E          api.Exception
	OutTradeNO string `json:"out_trade_no"` // 商户订单号
	TradeNO    string `json:"trade_no"`     // 支付宝交易号

}

// SetTags SetTags
func (resp *AlipayTradeCreateResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_create_response"
	exceptionTag = "alipay_trade_create_response"
	e = &resp.E
	return
}