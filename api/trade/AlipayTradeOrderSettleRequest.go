package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradeOrderSettleRequest alipay.trade.order.settle(统一收单交易结算接口)
// 用于在线下场景交易支付后，进行结算
type AlipayTradeOrderSettleRequest struct {
	api.PublicAppAuthToken
	OutRequestNO      string                       `json:"out_request_no"`     // 必选 结算请求流水号 开发者自行生成并保证唯一性
	TradeNO           string                       `json:"trade_no"`           // 必选 支付宝订单号
	RoyaltyParameters OpenAPIRoyaltyDetailInfoPojo `json:"royalty_parameters"` // 必选 分账明细信息
	OperatorID        string                       `json:"operator_id"`        // 可选 操作员id

}

// OpenAPIRoyaltyDetailInfoPojo OpenAPIRoyaltyDetailInfoPojo
type OpenAPIRoyaltyDetailInfoPojo struct {
	TransOut         string  `json:"trans_out"`         // 可选 分账支出方账户，类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	TransIn          string  `json:"trans_in"`          // 可选 分账收入方账户，类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	Amount           float64 `json:"amount"`            // 可选 分账的金额，单位为元
	AmountPercentage float64 `json:"amount_percentage"` // 可选 分账信息中分账百分比。取值范围为大于0，少于或等于100的整数。
	Desc             string  `json:"desc"`              // 可选 分账明细信息

}

// GetAPImethodName GetAPImethodName
func (requ *AlipayTradeOrderSettleRequest) GetAPImethodName() string {
	return "alipay.trade.fastpay.refund.query"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayTradeOrderSettleRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayTradeOrderSettleRequest) IsNeedBizContent() bool {
	return true
}