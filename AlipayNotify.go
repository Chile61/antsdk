package antsdk

// AlipayNotify 异步通知参数 https://docs.open.alipay.com/204/105301/
type AlipayNotify struct {
	NotifyTime        string `json:"notify_time"`         // 通知的发送时间。格式为yyyy-MM-dd HH:mm:ss
	NotifyType        string `json:"notify_type"`         // 通知的类型
	NotifyID          string `json:"notify_id"`           // 通知校验ID
	Charset           string `json:"charset"`             // (手机网站支付、App支付)编码格式，如utf-8、gbk、gb2312等
	Version           string `json:"version"`             // (手机网站支付、App支付)调用的接口版本，固定为：1.0
	SignType          string `json:"sign_type"`           // 签名算法类型，目前支持RSA
	Sign              string `json:"sign"`                // 请参考异步返回结果的验签(https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.FMGCDR&treeId=204&articleId=105301&docType=1#s6)
	AppID             string `json:"app_id"`              // 支付宝分配给开发者的应用Id
	AuthAppID         string `json:"auth_app_id"`         // 授权方的appid，由于本接口暂不开放第三方应用授权，因此auth_app_id=app_id
	TradeNo           string `json:"trade_no"`            // 支付宝交易凭证号
	OutTradeNo        string `json:"out_trade_no"`        // 原支付请求的商户订单号
	OutBizNo          string `json:"out_biz_no"`          // 商户业务ID，主要是退款通知中返回退款申请的流水号
	BuyerID           string `json:"buyer_id"`            // 买家支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字
	BuyerEmail        string `json:"buyer_email"`         // 买家支付宝账号，可以是Email或手机号码。
	TotalFee          string `json:"total_fee"`           // 该笔订单的总金额。请求时对应的参数，原样通知回来。
	BuyerLogonID      string `json:"buyer_logon_id"`      // 买家支付宝账号
	SellerID          string `json:"seller_id"`           // 卖家支付宝用户号
	SellerEmail       string `json:"seller_email"`        // 卖家支付宝账号
	TradeStatus       string `json:"trade_status"`        // 交易目前所处的状态，见交易状态说明(https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.FMGCDR&treeId=204&articleId=105301&docType=1#s1)
	TotalAmount       string `json:"total_amount"`        // 本次交易支付的订单金额，单位为人民币（元）
	ReceiptAmount     string `json:"receipt_amount"`      // 商家在交易中实际收到的款项，单位为元
	InvoiceAmount     string `json:"invoice_amount"`      // 用户在交易中支付的可开发票的金额
	BuyerPayAmount    string `json:"buyer_pay_amount"`    // 用户在交易中支付的金额
	PointAmount       string `json:"point_amount"`        // 使用集分宝支付的金额
	RefundFee         string `json:"refund_fee"`          // 退款通知中，返回总退款金额，单位为元，支持两位小数
	Subject           string `json:"subject"`             // 商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来
	Body              string `json:"body"`                // 该订单的备注、描述、明细等。对应请求时的body参数，原样通知回来
	GMTCreate         string `json:"gmt_create"`          // 该笔交易创建的时间。格式为yyyy-MM-dd HH:mm:ss
	GMTPayment        string `json:"gmt_payment"`         // 该笔交易的买家付款时间。格式为yyyy-MM-dd HH:mm:ss
	GMTRefund         string `json:"gmt_refund"`          // 该笔交易的退款时间。格式为yyyy-MM-dd HH:mm:ss.S
	GMTClose          string `json:"gmt_close"`           // 该笔交易结束时间。格式为yyyy-MM-dd HH:mm:ss
	FundBillList      string `json:"fund_bill_list"`      // 支付成功的各个渠道金额信息，详见资金明细信息说明(https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.FMGCDR&treeId=204&articleId=105301&docType=1#s3)
	PassbackParams    string `json:"passback_params"`     // 共回传参数，如果请求时传递了该参数，则返回给商户时会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝
	VoucherDetailList string `json:"voucher_detail_list"` // 本交易支付时所使用的所有优惠券信息，详见优惠券信息说明(https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.FMGCDR&treeId=204&articleId=105301&docType=1#s5)
	SendBackFee       string `json:"send_back_fee"`       // (当面付)商户实际退款给用户的金额，单位为元，支持两位小数
}
