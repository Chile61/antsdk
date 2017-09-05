package market

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 服务商完成订单内单个明细实施项
// 商户在订购插件后，服务商完成实施操作，无须授权，仅需检测当前订单属于当前操作服务商
type AlipayOpenServicemarketOrderItemCompleteRequest struct {
	api.IAlipayRequest
	TerminalType string                                                    `json:"terminal_type"`
	TerminalInfo string                                                    `json:"terminal_info"`
	ProdCode     string                                                    `json:"prod_code"`
	NotifyURL    string                                                    `json:"notify_url"`
	ReturnURL    string                                                    `json:"return_url"`
	BizContent   AlipayOpenServicemarketOrderItemCompleteRequestBizContent `json:"biz_content"`
}

type AlipayOpenServicemarketOrderItemCompleteRequestBizContent struct {
	CommodityOrderId string `json:"commodity_order_id"` // 订购服务插件订单号
	ShopId           string `json:"shop_id"`            // 订购插件选择的某一店铺ID
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetAPImethodName() string {
	return "alipay.open.servicemarket.order.item.complete"
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenServicemarketOrderItemCompleteRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	return txtParams
}
