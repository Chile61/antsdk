package data

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 自定义报表规则创建及更新接口
type KoubeiMarketingDataCustomreportSaveRequest struct {
	api.IAlipayRequest
	TerminalType string                                               `json:"terminal_type"`
	TerminalInfo string                                               `json:"terminal_info"`
	ProdCode     string                                               `json:"prod_code"`
	NotifyURL    string                                               `json:"notify_url"`
	ReturnURL    string                                               `json:"return_url"`
	BizContent   KoubeiMarketingDataCustomreportSaveRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportSaveRequestBizContent struct {
	ReportConditionInfo CustomReportCondition `json:"report_condition_info"`
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetAPImethodName() string {
	return "koubei.marketing.data.customreport.save"
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetAPIversion() string {
	return "1.0"
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingDataCustomreportSaveRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
