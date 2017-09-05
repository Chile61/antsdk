package data

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 自定义报表规则列表分页查询接口
// 分页查询自定义数据报表规则列表
type KoubeiMarketingDataCustomreportBatchqueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                                     `json:"terminal_type"`
	TerminalInfo string                                                     `json:"terminal_info"`
	ProdCode     string                                                     `json:"prod_code"`
	NotifyURL    string                                                     `json:"notify_url"`
	ReturnURL    string                                                     `json:"return_url"`
	BizContent   KoubeiMarketingDataCustomreportBatchqueryRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportBatchqueryRequestBizContent struct {
	PageNo   string `json:"page_no"`   // 当前页号，默认为1
	PageSize string `json:"page_size"` // 每页条目数，默认为20,最大为30
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetAPImethodName() string {
	return "koubei.marketing.data.customreport.batchquery"
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingDataCustomreportBatchqueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
