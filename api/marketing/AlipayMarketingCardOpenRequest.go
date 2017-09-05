package marketing

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 会员卡开卡
// 商户通过API接口，进行会员卡开卡。
type AlipayMarketingCardOpenRequest struct {
	api.IAlipayRequest
	TerminalType string                                   `json:"terminal_type"`
	TerminalInfo string                                   `json:"terminal_info"`
	ProdCode     string                                   `json:"prod_code"`
	NotifyURL    string                                   `json:"notify_url"`
	ReturnURL    string                                   `json:"return_url"`
	BizContent   AlipayMarketingCardOpenRequestBizContent `json:"biz_content"`
}

type AlipayMarketingCardOpenRequestBizContent struct {
	OutSerialNo    string         `json:"out_serial_no"`    // 外部商户流水号（商户需要确保唯一性控制，类似request_id唯一请求标识）
	CardTemplateId string         `json:"card_template_id"` // 支付宝分配的卡模板Id（卡模板创建接口返回的模板ID）
	CardUserInfo   CardUserInfo   `json:"card_user_info"`   // 发卡用户信息
	CardExtInfo    MerchantCard   `json:"card_ext_info"`    // 外部卡信息(biz_card_no无需填写)
	MemberExtInfo  MerchantMenber `json:"member_ext_info"`  // 商户会员信息
}

func (this *AlipayMarketingCardOpenRequest) GetAPImethodName() string {
	return "alipay.marketing.card.open"
}

func (this *AlipayMarketingCardOpenRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayMarketingCardOpenRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayMarketingCardOpenRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayMarketingCardOpenRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayMarketingCardOpenRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayMarketingCardOpenRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayMarketingCardOpenRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayMarketingCardOpenRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
