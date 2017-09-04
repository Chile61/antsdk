package coupon

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 集点查询
// 查询用户集点
type KoubeiMarketingToolPointsQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                          `json:"terminal_type"`
	TerminalInfo string                                          `json:"terminal_info"`
	ProdCode     string                                          `json:"prod_code"`
	NotifyURL    string                                          `json:"notify_url"`
	ReturnURL    string                                          `json:"return_url"`
	BizContent   KoubeiMarketingToolPointsQueryRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingToolPointsQueryRequestBizContent struct {
	ActivityAccount string `json:"activity_account"` // 活动积分帐户
	UserId          string `json:"user_id"`          // 用户ID
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetApiMethodName() string {
	return "koubei.marketing.tool.points.query"
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetApiVersion() string {
	return "1.0"
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingToolPointsQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingToolPointsQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
