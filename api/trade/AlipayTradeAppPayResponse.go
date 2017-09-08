package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradeAppPayResponse AlipayTradeAppPayResponse
type AlipayTradeAppPayResponse struct {
	E api.Exception
}

// SetTags SetTags
func (resp *AlipayTradeAppPayResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_close_response"
	exceptionTag = "alipay_trade_close_response"
	e = &resp.E
	return
}