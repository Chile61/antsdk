package shop

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 查询商户的门店编号列表
// 系统商通过该接口可以查询对应APPID下所有的门店编号（支付宝口碑门店编号）
type AlipayOfflineMarketShopBatchqueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                             `json:"terminal_type"`
	TerminalInfo string                                             `json:"terminal_info"`
	ProdCode     string                                             `json:"prod_code"`
	NotifyURL    string                                             `json:"notify_url"`
	ReturnURL    string                                             `json:"return_url"`
	BizContent   AlipayOfflineMarketShopBatchqueryRequestBizContent `json:"biz_content"`
}

type AlipayOfflineMarketShopBatchqueryRequestBizContent struct {
	PageNo string `json:"page_no"` // 页码，第一页传入"1"，默认500个结果为一页
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetAPImethodName() string {
	return "alipay.offline.market.shop.batchquery"
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOfflineMarketShopBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
