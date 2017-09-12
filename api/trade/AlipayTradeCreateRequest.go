package trade

import "../../api"

// AlipayTradeCreateRequest alipay.trade.create(统一收单交易创建接口)
// 商户通过该接口进行交易的创建下单
type AlipayTradeCreateRequest struct {
	api.PublicAppAuthToken
	api.PublicNotifyURL
	OutTradeNO         string        `json:"out_trade_no"`        // 必选  商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	SellerID           string        `json:"seller_id"`           // 可选  如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount        string        `json:"total_amount"`        // 可选  订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入【可打折金额】和【不可打折金额】，该参数可以不用传入； 如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
	DiscountableAmount string        `json:"discountable_amount"` // 可选  参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	Subject            string        `json:"subject"`             // 必选  订单标题
	Body               string        `json:"body"`                // 可选  订单描述
	GoodsDetailList    []GoodsDetail `json:"GoodsDetail"`         // 可选  订单包含的商品列表信息，Json格式，其它说明详见商品明细说明
	OperatorID         string        `json:"operator_id"`         // 可选  商户操作员编号
	StoreID            string        `json:"store_id"`            // 可选  商户门店编号
	TerminalID         string        `json:"terminal_id"`         // 可选  商户机具终端编号
	ExtendParamsList   ExtendParams  `json:"extend_params"`       // 可选  业务扩展参数
	TimeoutExpress     string        `json:"timeout_express"`     // 可选  该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m

}

// GetAPImethodName GetAPImethodName
func (requ *AlipayTradeCreateRequest) GetAPImethodName() string {
	return "alipay.trade.create"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayTradeCreateRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayTradeCreateRequest) IsNeedBizContent() bool {
	return true
}
