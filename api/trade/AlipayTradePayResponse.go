package trade

import "../../api"

// AlipayTradePayResponse AlipayTradePayResponse
type AlipayTradePayResponse struct {
	E                   api.Exception
	TradeNO             string          `json:"trade_no"`              // 支付宝交易号
	OutTradeNO          string          `json:"out_trade_no"`          // 商户订单号
	BuyerLogonID        string          `json:"buyer_logon_id"`        // 买家支付宝账号
	TotalAmount         string          `json:"total_amount"`          // 交易金额
	ReceiptAmount       string          `json:"receipt_amount"`        // 实收金额
	BuyerPayAmount      string          `json:"buyer_pay_amount"`      // 买家付款的金额
	PointAmount         string          `json:"point_amount"`          // 使用积分宝付款的金额
	InvoiceAmount       string          `json:"invoice_amount"`        // 交易中可给用户开具发票的金额
	GMTPayment          string          `json:"gmt_payment"`           // 交易支付时间
	FundBillList        []FundBill      `json:"fund_bill_list"`        // 交易支付使用的资金渠道
	CardBalance         string          `json:"card_balance"`          // 支付宝卡余额
	StoreName           string          `json:"store_name"`            // 发生支付交易的商户门店名称
	BuyerUserID         string          `json:"buyer_user_id"`         // 买家在支付宝的用户id
	DiscountGoodsDetail string          `json:"discount_goods_detail"` // 本次交易支付所使用的单品券优惠的商品优惠信息
	VoucherDetailList   []VoucherDetail `json:"VoucherDetail"`         // 本交易支付时使用的所有优惠券信息

}

// VoucherDetail VoucherDetail
type VoucherDetail struct {
	ID                         string `json:"id"`                           // 券id
	Name                       string `json:"name"`                         // 券名称
	Type                       string `json:"type"`                         // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券 ALIPAY_DISCOUNT_VOUCHER - 折扣券 ALIPAY_ITEM_VOUCHER - 单品优惠 注：不排除将来新增其他类型的可能，商家接入时注意兼容性避免硬编码
	Amount                     string `json:"amount"`                       // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute         string `json:"merchant_contribute"`          // 商家出资（特指发起交易的商家出资金额）
	OtherContribute            string `json:"other_contribute"`             // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo                       string `json:"memo"`                         // 优惠券备注信息
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute"`    // 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时用户实际付款的金额
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute"` // 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时商户优惠的金额
	PurchaseAntContribute      string `json:"purchase_ant_contribute"`      // 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时平台优惠的金额
}

// SetTags SetTags
func (resp *AlipayTradePayResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_pay_response"
	exceptionTag = "alipay_trade_pay_response"
	e = &resp.E
	return
}
