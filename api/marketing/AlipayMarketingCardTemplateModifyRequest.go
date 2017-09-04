package marketing

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 会员卡模板修改
type AlipayMarketingCardTemplateModifyRequest struct {
	api.IAlipayRequest
	TerminalType string                                             `json:"terminal_type"`
	TerminalInfo string                                             `json:"terminal_info"`
	ProdCode     string                                             `json:"prod_code"`
	NotifyURL    string                                             `json:"notify_url"`
	ReturnURL    string                                             `json:"return_url"`
	BizContent   AlipayMarketingCardTemplateModifyRequestBizContent `json:"biz_content"`
}

type AlipayMarketingCardTemplateModifyRequestBizContent struct {
	RequestId           string                     `json:"request_id"`
	TemplateId          string                     `json:"template_id"`
	BizNoPrefix         string                     `json:"biz_no_prefix"`
	WriteOffType        string                     `json:"write_off_type"`
	TemplateStyleInfo   TemplateStyleInfoDTO       `json:"template_style_info"`
	TemplateBenefitInfo []TemplateBenefitInfoDTO   `json:"template_benefit_info"`
	ColumnInfoList      []TemplateColumnInfoDTO    `json:"column_info_list"`
	FieldRuleList       []TemplateFieldRuleDTO     `json:"field_rule_list"`
	ShopIds             []string                   `json:"shop_ids"`
	OpenCardConf        TemplateOpenCardConfDTO    `json:"open_card_conf"`
	PubChannels         []PubChannelDTO            `json:"pub_channels"`
	CardLevelConf       []TemplateCardLevelConfDTO `json:"card_level_conf"`
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetApiMethodName() string {
	return "alipay.marketing.card.template.modify"
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayMarketingCardTemplateModifyRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayMarketingCardTemplateModifyRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
