package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradeCloseRequest alipay.trade.close(统一收单交易关闭接口)
// 用于交易创建后，用户在一定时间内未进行支付，可调用该接口直接将未付款的交易进行关闭。
type AlipayTradeCloseRequest struct {
	api.PublicAppAuthToken
	api.PublicNotifyURL
	TradeNO    string `json:"trade_no"`     // 特殊可选 该交易在支付宝系统中的交易流水号。最短 16 位，最长 64 位。和out_trade_no不能同时为空，如果同时传了 out_trade_no和 trade_no，则以 trade_no为准。
	OutTradeNO string `json:"out_trade_no"` // 特殊可选  订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	OperatorID string `json:"operator_id"`  // 可选 卖家端自定义的的操作员 ID

}

// GetAPImethodName GetAPImethodName
func (requ *AlipayTradeCloseRequest) GetAPImethodName() string {
	return "alipay.trade.close"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayTradeCloseRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayTradeCloseRequest) IsNeedBizContent() bool {
	return true
}