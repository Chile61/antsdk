package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradeCancelRequest alipay.trade.cancel(统一收单交易撤销接口)
// 支付交易返回失败或支付系统超时，调用该接口撤销交易。
// 如果此订单用户支付失败，支付宝系统会将此订单关闭；
// 如果用户支付成功，支付宝系统会将此订单资金退还给用户。
// 注意：只有发生支付系统超时或者支付结果未知时可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。
// 提交支付交易后调用【查询订单API】，没有明确的支付结果再调用【撤销订单API】。
type AlipayTradeCancelRequest struct {
	api.PublicAppAuthToken
	TradeNO    string `json:"trade_no"`     // 特殊可选 原支付请求的商户订单号,和支付宝交易号不能同时为空
	OutTradeNO string `json:"out_trade_no"` // 特殊可选		支付宝交易号，和商户订单号不能同时为空
}

// GetAPImethodName GetAPImethodName
func (requ *AlipayTradeCancelRequest) GetAPImethodName() string {
	return "alipay.trade.cancel"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayTradeCancelRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayTradeCancelRequest) IsNeedBizContent() bool {
	return true
}