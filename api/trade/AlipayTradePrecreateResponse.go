package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradePrecreateResponse AlipayTradePrecreateResponse
type AlipayTradePrecreateResponse struct {
	E          api.Exception
	OutTradeNO string `json:"out_trade_no"` // 商户订单号
	QRCode     string `json:"qr_code"`      // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码

}

// SetTags SetTags
func (resp *AlipayTradePrecreateResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_precreate_response"
	exceptionTag = "alipay_trade_precreate_response"
	e = &resp.E
	return
}