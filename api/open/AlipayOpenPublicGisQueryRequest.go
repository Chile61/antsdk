package open

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 获取用户地理位置
// 查询该用户当前的位置信息，可以根据用户位置完成不同的业务逻辑。用户须关注该服务窗且允许上传地理位置，开发者才能通过此接口获取该用户的地理位置。用户进入具体服务窗时会定时上传地理位置，退出时停止上传，开发者调用此接口获取到的是用户最后一次上传的地理位置。
type AlipayOpenPublicGisQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                    `json:"terminal_type"`
	TerminalInfo string                                    `json:"terminal_info"`
	ProdCode     string                                    `json:"prod_code"`
	NotifyURL    string                                    `json:"notify_url"`
	ReturnURL    string                                    `json:"return_url"`
	BizContent   AlipayOpenPublicGisQueryRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicGisQueryRequestBizContent struct {
	UserId string `json:"user_id"` // 该用户的userId
}

func (this *AlipayOpenPublicGisQueryRequest) GetAPImethodName() string {
	return "alipay.open.public.gis.query"
}

func (this *AlipayOpenPublicGisQueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayOpenPublicGisQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenPublicGisQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenPublicGisQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenPublicGisQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenPublicGisQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenPublicGisQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenPublicGisQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
