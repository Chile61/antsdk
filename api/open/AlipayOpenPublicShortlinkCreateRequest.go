package open

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 服务窗短链自主生成接口
// 商户通过本接口生成带自有场景标识的短链接
type AlipayOpenPublicShortlinkCreateRequest struct {
	api.IAlipayRequest
	TerminalType string                                           `json:"terminal_type"`
	TerminalInfo string                                           `json:"terminal_info"`
	ProdCode     string                                           `json:"prod_code"`
	NotifyURL    string                                           `json:"notify_url"`
	ReturnURL    string                                           `json:"return_url"`
	BizContent   AlipayOpenPublicShortlinkCreateRequestBizContent `json:"biz_content"`
}

type AlipayOpenPublicShortlinkCreateRequestBizContent struct {
	SceneId string `json:"scene_id"` // 短链接对应的场景ID，该ID由商户自己定义
	Remark  string `json:"remark"`   // 对于场景ID的描述，商户自己定义
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetApiMethodName() string {
	return "alipay.open.public.shortlink.create"
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetApiVersion() string {
	return "1.0"
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayOpenPublicShortlinkCreateRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayOpenPublicShortlinkCreateRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
