package open

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 单发模板消息
type AlipayOpenPublicMessageSingleSendRequest struct {
	api.IAlipayRequest
	TerminalType string                                             `json:"terminal_type"`
	TerminalInfo string                                             `json:"terminal_info"`
	ProdCode     string                                             `json:"prod_code"`
	NotifyURL    string                                             `json:"notify_url"`
	ReturnURL    string                                             `json:"return_url"`
	BizContent   AlipayOpenPublicMessageSingleSendRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicMessageSingleSendRequestBizContent struct {
	ToUserId string   `json:"to_user_id"`
	Template Template `json:"template"`
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetApiMethodName() string {
	return "alipay.open.public.message.single.send"
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenPublicMessageSingleSendRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenPublicMessageSingleSendRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
