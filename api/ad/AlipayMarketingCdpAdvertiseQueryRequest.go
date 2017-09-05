package ad

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 查询广告接口
// 开发者可查询广告内容和状态。
type AlipayMarketingCdpAdvertiseQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                            `json:"terminal_type"`
	TerminalInfo string                                            `json:"terminal_info"`
	ProdCode     string                                            `json:"prod_code"`
	NotifyURL    string                                            `json:"notify_url"`
	ReturnURL    string                                            `json:"return_url"`
	BizContent   AlipayMarketingCdpAdvertiseQueryRequestBizContent `json:"biz_content"`
}

type AlipayMarketingCdpAdvertiseQueryRequestBizContent struct {
	AdId string `json:"ad_id"` // 广告Id，唯一标识一条广告
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetAPImethodName() string {
	return "alipay.marketing.cdp.advertise.query"
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayMarketingCdpAdvertiseQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
