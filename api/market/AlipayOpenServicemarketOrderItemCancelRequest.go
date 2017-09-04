package market

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 服务订单明细实施项单项取消
// 当服务商订购服务区域较多时，可能部分区域不能实施，可以取消实施
type AlipayOpenServicemarketOrderItemCancelRequest struct {
	api.IAlipayRequest
	TerminalType string                                                  `json:"terminal_type"`
	TerminalInfo string                                                  `json:"terminal_info"`
	ProdCode     string                                                  `json:"prod_code"`
	NotifyURL    string                                                  `json:"notify_url"`
	ReturnURL    string                                                  `json:"return_url"`
	BizContent   AlipayOpenServicemarketOrderItemCancelRequestBizContent `json:"biz_content"`
}

type AlipayOpenServicemarketOrderItemCancelRequestBizContent struct {
	CommodityOrderId string `json:"commodity_order_id"` // 订购服务订单ID
	ShopId           string `json:"shop_id"`            // 订购服务门店ID
	CancelReason     string `json:"cancel_reason"`      // 当前门店区域不支持实施
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetApiMethodName() string {
	return "alipay.open.servicemarket.order.item.cancel"
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenServicemarketOrderItemCancelRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	return txtParams
}
