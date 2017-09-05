package marketing

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 会员卡更新
type AlipayMarketingCardUpdateRequest struct {
	api.IAlipayRequest
	TerminalType string                                     `json:"terminal_type"`
	TerminalInfo string                                     `json:"terminal_info"`
	ProdCode     string                                     `json:"prod_code"`
	NotifyURL    string                                     `json:"notify_url"`
	ReturnURL    string                                     `json:"return_url"`
	BizContent   AlipayMarketingCardUpdateRequestBizContent `json:"biz_content"`
}

type AlipayMarketingCardUpdateRequestBizContent struct {
	TargetCardNo     string       `json:"target_card_no"`      // 支付宝业务卡号，开卡接口中返回获取
	TargetCardNoType string       `json:"target_card_no_type"` // 卡号ID类型 BIZ_CARD：支付宝业务卡号
	OccurTime        string       `json:"occur_time"`          // 标识业务发生的时间
	CardInfo         MerchantCard `json:"card_info"`           // 需要修改的最新卡信息
	ExtInfo          string       `json:"ext_info"`            // 扩展信息(暂时无用)
}

func (this *AlipayMarketingCardUpdateRequest) GetAPImethodName() string {
	return "alipay.marketing.card.update"
}

func (this *AlipayMarketingCardUpdateRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayMarketingCardUpdateRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayMarketingCardUpdateRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayMarketingCardUpdateRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayMarketingCardUpdateRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayMarketingCardUpdateRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayMarketingCardUpdateRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayMarketingCardUpdateRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
