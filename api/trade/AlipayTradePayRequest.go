package trade

import "github.com/vanishs/antsdk/api"

// AlipayTradePayRequest alipay.trade.pay(统一收单交易支付接口)
// 收银员使用扫码设备读取用户手机支付宝“付款码”/声波获取设备（如麦克风）读取用户手机支付宝的声波信息后，将二维码或条码信息/声波信息通过本接口上送至支付宝发起支付。
type AlipayTradePayRequest struct {
	api.PublicAppAuthToken
	api.PublicNotifyURL
	OutTradeNO         string        `json:"out_trade_no"`        // 必选  商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	Scene              string        `json:"scene"`               // 必选  订	支付场景 条码支付，取值：bar_code 声波支付，取值：wave_code
	AuthCode           string        `json:"auth_code"`           // 必选  支付授权码，25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准
	ProductCode        string        `json:"product_code"`        // 可选  销售产品码
	Subject            string        `json:"subject"`             // 必选  订单标题
	BuyerID            string        `json:"buyer_id"`            // 可选  买家的支付宝用户id，如果为空，会从传入了码值信息中获取买家ID
	SellerID           string        `json:"seller_id"`           // 可选  如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount        string        `json:"total_amount"`        // 可选  订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入【可打折金额】和【不可打折金额】，该参数可以不用传入； 如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
	DiscountableAmount string        `json:"discountable_amount"` // 可选  参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	Body               string        `json:"body"`                // 可选  订单描述
	GoodsDetailList    []GoodsDetail `json:"GoodsDetail"`         // 可选  订单包含的商品列表信息，Json格式，其它说明详见商品明细说明
	OperatorID         string        `json:"operator_id"`         // 可选  商户操作员编号
	StoreID            string        `json:"store_id"`            // 可选  商户门店编号
	TerminalID         string        `json:"terminal_id"`         // 可选  商户机具终端编号
	ExtendParamsList   ExtendParams  `json:"extend_params"`       // 可选  业务扩展参数
	TimeoutExpress     string        `json:"timeout_express"`     // 可选  该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m

}

// GoodsDetail GoodsDetail
type GoodsDetail struct {
	GoodsID       string `json:"goods_id"`       // 必填  商品的编号
	GoodsName     string `json:"goods_name"`     // 必填  商品名称
	Quantity      string `json:"quantity"`       // 必填  商品数量
	Price         string `json:"price"`          // 必填  商品单价，单位为元
	GoodsCategory string `json:"goods_category"` // 可选  商品类目
	Body          string `json:"body"`           // 可选  商品描述信息
	ShowURL       string `json:"show_url"`       // 可选  商品的展示地址

}

// ExtendParams ExtendParams
type ExtendParams struct {
	SysServiceProviderID string `json:"sys_service_provider_id"` // 可选  系统商编号 该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID

}

// GetAPImethodName GetAPImethodName
func (requ *AlipayTradePayRequest) GetAPImethodName() string {
	return "alipay.trade.pay"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayTradePayRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayTradePayRequest) IsNeedBizContent() bool {
	return true
}