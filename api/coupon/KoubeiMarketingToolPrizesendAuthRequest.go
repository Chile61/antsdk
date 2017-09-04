package coupon

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 发券授权
// 返回token授权商户给用户发奖
type KoubeiMarketingToolPrizesendAuthRequest struct {
	api.IAlipayRequest
	TerminalType string                                            `json:"terminal_type"`
	TerminalInfo string                                            `json:"terminal_info"`
	ProdCode     string                                            `json:"prod_code"`
	NotifyURL    string                                            `json:"notify_url"`
	ReturnURL    string                                            `json:"return_url"`
	BizContent   KoubeiMarketingToolPrizesendAuthRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingToolPrizesendAuthRequestBizContent struct {
	ReqId   string `json:"req_id"`   // 外部流水号，保证业务幂等性
	PrizeId string `json:"prize_id"` // 奖品ID
	UserId  string `json:"user_id"`  // 发奖用户ID
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetApiMethodName() string {
	return "koubei.marketing.tool.prizesend.auth"
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetApiVersion() string {
	return "1.0"
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingToolPrizesendAuthRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
