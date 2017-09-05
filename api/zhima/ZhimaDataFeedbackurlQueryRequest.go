package zhima

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 获取数据反馈模板
// isv商户获取数据反馈模板
type ZhimaDataFeedbackurlQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                     `json:"terminal_type"`
	TerminalInfo string                                     `json:"terminal_info"`
	ProdCode     string                                     `json:"prod_code"`
	NotifyURL    string                                     `json:"notify_url"`
	ReturnURL    string                                     `json:"return_url"`
	BizContent   ZhimaDataFeedbackurlQueryRequestBizContent `json:"biz_content"`
}

type ZhimaDataFeedbackurlQueryRequestBizContent struct {
	MerchantId string `json:"merchant_id"`
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetAPImethodName() string {
	return "zhima.data.feedbackurl.query"
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *ZhimaDataFeedbackurlQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *ZhimaDataFeedbackurlQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
