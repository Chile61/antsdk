package data

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 自定义报表规则删除接口
// 删除自定义数据报表规则
type KoubeiMarketingDataCustomreportDeleteRequest struct {
	api.IAlipayRequest
	TerminalType string                                                 `json:"terminal_type"`
	TerminalInfo string                                                 `json:"terminal_info"`
	ProdCode     string                                                 `json:"prod_code"`
	NotifyURL    string                                                 `json:"notify_url"`
	ReturnURL    string                                                 `json:"return_url"`
	BizContent   KoubeiMarketingDataCustomreportDeleteRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportDeleteRequestBizContent struct {
	ConditionKey string `json:"condition_key"` // 自定义报表规则的KEY
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetAPImethodName() string {
	return "koubei.marketing.data.customreport.delete"
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetAPIversion() string {
	return "1.0"
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingDataCustomreportDeleteRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
