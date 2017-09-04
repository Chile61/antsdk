package zhima

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 芝麻认证开始认证
// 在商户完成认证初始化后,可以开始认证了
type ZhimaCustomerCertificationCertifyRequest struct {
	api.IAlipayRequest
	TerminalType string                                             `json:"terminal_type"`
	TerminalInfo string                                             `json:"terminal_info"`
	ProdCode     string                                             `json:"prod_code"`
	NotifyURL    string                                             `json:"notify_url"`
	ReturnURL    string                                             `json:"return_url"`
	BizContent   ZhimaCustomerCertificationCertifyRequestBizContent `json:"biz_content"`
}

type ZhimaCustomerCertificationCertifyRequestBizContent struct {
	BizNo string `json:"biz_no"` // 一次认证的唯一标识,在完成芝麻认证初始化后可以获取
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetApiMethodName() string {
	return "zhima.customer.certification.certify"
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetApiVersion() string {
	return "1.0"
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *ZhimaCustomerCertificationCertifyRequest) IsNeedEncrypt() bool {
	return false
}

func (this *ZhimaCustomerCertificationCertifyRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
