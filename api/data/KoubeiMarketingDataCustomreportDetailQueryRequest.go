package data

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 自定义报表规则详情查询接口
// 查询自定义数据报表规则详情
type KoubeiMarketingDataCustomreportDetailQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                                      `json:"terminal_type"`
	TerminalInfo string                                                      `json:"terminal_info"`
	ProdCode     string                                                      `json:"prod_code"`
	NotifyURL    string                                                      `json:"notify_url"`
	ReturnURL    string                                                      `json:"return_url"`
	BizContent   KoubeiMarketingDataCustomreportDetailQueryRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingDataCustomreportDetailQueryRequestBizContent struct {
	ConditionKey string `json:"condition_key"` // 自定义报表的规则KEY
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetApiMethodName() string {
	return "koubei.marketing.data.customreport.detail.query"
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetApiVersion() string {
	return "1.0"
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingDataCustomreportDetailQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
