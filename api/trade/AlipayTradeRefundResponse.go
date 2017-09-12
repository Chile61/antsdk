package trade

import "../../api"

// AlipayTradeRefundResponse AlipayTradeRefundResponse
type AlipayTradeRefundResponse struct {
	E                    api.Exception
	TradeNO              string     `json:"trade_no"`                // 支付宝交易号
	OutTradeNO           string     `json:"out_trade_no"`            // 商户订单号
	BuyerLogonID         string     `json:"buyer_logon_id"`          // 用户的登录id
	FundChange           string     `json:"fund_change"`             // 本次退款是否发生了资金变化 示例值Y
	RefundFee            string     `json:"refund_fee"`              // 退款总金额
	GMTRefundPay         string     `json:"gmt_refund_pay"`          // 退款支付时间
	RefundDetailItemList []FundBill `json:"refund_detail_item_list"` // 退款使用的资金渠道
	StoreName            string     `json:"store_name"`              // 交易在支付时候的门店名称
	BuyerUserID          string     `json:"buyer_user_id"`           // 买家在支付宝的用户id

}

// FundBill FundBill
type FundBill struct {
	FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表 https://doc.open.alipay.com/doc2/detail?treeId=26&articleId=103259&docType=1
	Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
	RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
	FundType    string `json:"fund_type"`    // 渠道所使用的资金类型,目前只在资金渠道(fund_channel)是银行卡渠道(BANKCARD)的情况下才返回该信息(DEBIT_CARD:借记卡,CREDIT_CARD:信用卡,MIXED_CARD:借贷合一卡)
}

// SetTags SetTags
func (resp *AlipayTradeRefundResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_refund_response"
	exceptionTag = "alipay_trade_refund_response"
	e = &resp.E
	return
}
