package open

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 创建菜单
type AlipayOpenPublicMenuCreateRequest struct {
	api.IAlipayRequest
	TerminalType string                                      `json:"terminal_type"`
	TerminalInfo string                                      `json:"terminal_info"`
	ProdCode     string                                      `json:"prod_code"`
	NotifyURL    string                                      `json:"notify_url"`
	ReturnURL    string                                      `json:"return_url"`
	BizContent   AlipayOpenPublicMenuCreateRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicMenuCreateRequestBizContent struct {
	Button []ButtonObject `json:"button"` // 一级菜单数组，个数应为1~4个
}

func (this *AlipayOpenPublicMenuCreateRequest) GetAPImethodName() string {
	return "alipay.open.public.menu.create"
}

func (this *AlipayOpenPublicMenuCreateRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenPublicMenuCreateRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenPublicMenuCreateRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenPublicMenuCreateRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenPublicMenuCreateRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenPublicMenuCreateRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenPublicMenuCreateRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenPublicMenuCreateRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
