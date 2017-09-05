package data

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 营销活动人群组规则列表分页查询接口
// 口碑商户人群组列表查询接口
type KoubeiMarketingCampaignCrowdBatchqueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                                  `json:"terminal_type"`
	TerminalInfo string                                                  `json:"terminal_info"`
	ProdCode     string                                                  `json:"prod_code"`
	NotifyURL    string                                                  `json:"notify_url"`
	ReturnURL    string                                                  `json:"return_url"`
	BizContent   KoubeiMarketingCampaignCrowdBatchqueryRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingCampaignCrowdBatchqueryRequestBizContent struct {
	Name       string `json:"name"`        // 人群名称
	PageNumber string `json:"page_number"` // 分页页码，从1开始计数,如果不填写默认为1
	PageSize   string `json:"page_size"`   // 分页大小，每页显示的数目，如果不填写默认为20
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetAPImethodName() string {
	return "koubei.marketing.campaign.crowd.batchquery"
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingCampaignCrowdBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
