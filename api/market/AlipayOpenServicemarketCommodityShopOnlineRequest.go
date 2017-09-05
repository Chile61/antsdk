package market

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 门店插件上架操作
// 本接口用于上架商户门店订购的服务。
type AlipayOpenServicemarketCommodityShopOnlineRequest struct {
	api.IAlipayRequest
	TerminalType string                                                      `json:"terminal_type"`
	TerminalInfo string                                                      `json:"terminal_info"`
	ProdCode     string                                                      `json:"prod_code"`
	NotifyURL    string                                                      `json:"notify_url"`
	ReturnURL    string                                                      `json:"return_url"`
	BizContent   AlipayOpenServicemarketCommodityShopOnlineRequestBizContent `json:"biz_content"`
}

type AlipayOpenServicemarketCommodityShopOnlineRequestBizContent struct {
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetAPImethodName() string {
	return "alipay.open.servicemarket.commodity.shop.online"
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenServicemarketCommodityShopOnlineRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	return txtParams
}
