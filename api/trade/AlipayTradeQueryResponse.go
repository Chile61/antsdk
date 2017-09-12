package trade

import "../../api"

// AlipayTradeQueryResponse AlipayTradeQueryResponse
type AlipayTradeQueryResponse struct {
	E              api.Exception
	TradeNO        string     `json:"trade_no"`         // 支付宝交易号
	OutTradeNO     string     `json:"out_trade_no"`     // 创建交易传入的商户订单号
	BuyerLogonID   string     `json:"buyer_logon_id"`   // 买家支付宝账号
	TradeStatus    string     `json:"trade_status"`     // 交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
	TotalAmount    string     `json:"total_amount"`     // 交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount
	ReceiptAmount  string     `json:"receipt_amount"`   // 	实收金额，单位为元，两位小数。该金额为本笔交易，商户账户能够实际收到的金额
	BuyerPayAmount string     `json:"buyer_pay_amount"` // 买家实付金额，单位为元，两位小数。该金额代表该笔交易买家实际支付的金额，不包含商户折扣等金额
	PointAmount    string     `json:"point_amount"`     // 	积分支付的金额，单位为元，两位小数。该金额代表该笔交易中用户使用积分支付的金额，比如集分宝或者支付宝实时优惠等
	InvoiceAmount  string     `json:"invoice_amount"`   // 交易中用户支付的可开具发票的金额，单位为元，两位小数。该金额代表该笔交易中可以给用户开具发票的金额
	SendPayDate    string     `json:"send_pay_date"`    // 	本次交易打款给卖家的时间
	StoreID        string     `json:"store_id"`         // 商户门店编号
	TerminalID     string     `json:"terminal_id"`      // 商户机具终端编号
	FundBillList   []FundBill `json:"fund_bill_list"`   // 交易支付使用的资金渠道
	StoreName      string     `json:"store_name"`       // 请求交易支付中的商户店铺的名称
	BuyerUserID    string     `json:"buyer_user_id"`    // 买家在支付宝的用户id

}

// SetTags SetTags
func (resp *AlipayTradeQueryResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_trade_query_response"
	exceptionTag = "alipay_trade_query_response"
	e = &resp.E
	return
}
