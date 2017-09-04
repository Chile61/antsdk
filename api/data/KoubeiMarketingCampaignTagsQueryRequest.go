package data

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

type KoubeiMarketingCampaignTagsQueryRequest struct {
	api.IAlipayRequest
	TerminalType string `json:"terminal_type"`
	TerminalInfo string `json:"terminal_info"`
	ProdCode     string `json:"prod_code"`
	NotifyURL    string `json:"notify_url"`
	ReturnURL    string `json:"return_url"`
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetApiMethodName() string {
	return "koubei.marketing.campaign.tags.query"
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetApiVersion() string {
	return "1.0"
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingCampaignTagsQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	return txtParams
}
