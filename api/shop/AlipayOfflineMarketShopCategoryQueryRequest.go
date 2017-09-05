package shop

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 门店类目配置查询接口
// 用于查询可用于开店的类目，以及类目上的配置约束信息
type AlipayOfflineMarketShopCategoryQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                                `json:"terminal_type"`
	TerminalInfo string                                                `json:"terminal_info"`
	ProdCode     string                                                `json:"prod_code"`
	NotifyURL    string                                                `json:"notify_url"`
	ReturnURL    string                                                `json:"return_url"`
	BizContent   AlipayOfflineMarketShopCategoryQueryRequestBizContent `json:"biz_content"`
}

type AlipayOfflineMarketShopCategoryQueryRequestBizContent struct {
	CategoryId string `json:"category_id"` // 类目ID，如果为空则查询全部类目。
	OPRole     string `json:"op_role"`     // 表示接口业务的调用方身份,默认不填标识为ISV。
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetAPImethodName() string {
	return "alipay.offline.market.shop.category.query"
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOfflineMarketShopCategoryQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
