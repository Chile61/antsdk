package open

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 更新菜单
// 通过POST一个特定结构体，实现支付宝钱包客户端的公众账号更新自定义菜单。每一次的更新是针对全部自定义菜单的更新。
type AlipayOpenPublicMenuModifyRequest struct {
	api.IAlipayRequest
	TerminalType string                                      `json:"terminal_type"`
	TerminalInfo string                                      `json:"terminal_info"`
	ProdCode     string                                      `json:"prod_code"`
	NotifyURL    string                                      `json:"notify_url"`
	ReturnURL    string                                      `json:"return_url"`
	BizContent   AlipayOpenPublicMenuModifyRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicMenuModifyRequestBizContent struct {
	Button []ButtonObject `json:"button"` // 一级菜单数组，个数应为1~4个
}

func (this *AlipayOpenPublicMenuModifyRequest) GetAPImethodName() string {
	return "alipay.open.public.menu.modify"
}

func (this *AlipayOpenPublicMenuModifyRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenPublicMenuModifyRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenPublicMenuModifyRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenPublicMenuModifyRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenPublicMenuModifyRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenPublicMenuModifyRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenPublicMenuModifyRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenPublicMenuModifyRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
