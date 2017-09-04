package marketing

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 会员卡模板查询接口
type AlipayMarketingCardTemplateQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                            `json:"terminal_type"`
	TerminalInfo string                                            `json:"terminal_info"`
	ProdCode     string                                            `json:"prod_code"`
	NotifyURL    string                                            `json:"notify_url"`
	ReturnURL    string                                            `json:"return_url"`
	BizContent   AlipayMarketingCardTemplateQueryRequestBizContent `json:"biz_content"`
}

type AlipayMarketingCardTemplateQueryRequestBizContent struct {
	TemplateId string `json:"template_id"` // 支付宝卡模板ID（模板创建接口返回的支付宝端模板ID）
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetApiMethodName() string {
	return "alipay.marketing.card.template.query"
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayMarketingCardTemplateQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayMarketingCardTemplateQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
