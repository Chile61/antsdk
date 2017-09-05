package coupon

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 更新卡集点
// 更新集点卡上的集点
type KoubeiMarketingToolPointsUpdateRequest struct {
	api.IAlipayRequest
	TerminalType string                                           `json:"terminal_type"`
	TerminalInfo string                                           `json:"terminal_info"`
	ProdCode     string                                           `json:"prod_code"`
	NotifyURL    string                                           `json:"notify_url"`
	ReturnURL    string                                           `json:"return_url"`
	BizContent   KoubeiMarketingToolPointsUpdateRequestBizContent `json:"biz_content"`
}

type KoubeiMarketingToolPointsUpdateRequestBizContent struct {
	ReqId           string `json:"req_id"`           // 外部流水号, 由开发者提供, 用于控制业务幂等
	UserId          string `json:"user_id"`          // 用户ID, 开发者通过用户信息授权产品获取
	ActivityAccount string `json:"activity_account"` // 活动集点帐户ID, 开发者通过查询集点活动详情获取
	TransType       string `json:"trans_type"`       // 集点交易类型，目前支持: DEPOSIT，增加集点 FREEZE，冻结集点 COMMIT，提交冻结集点 CANCEL，取消冻结集点 CONSUME, 消费集点
	TransAmount     string `json:"trans_amount"`     // 集点交易数量，必须为正整数字符串
	ShopId          string `json:"shop_id"`          // 门店ID，集点交易类型为DEPOSIT时填写
	BizNo           string `json:"biz_no"`           // 业务流水号，集点交易类型为 DEPOSIT, 传入支付交易号; CANCEL／COMMIT, 传入冻结集点的集点流水号; CONSUME／FREEZE, 不允许传入biz_no;
	Memo            string `json:"memo"`             // 交易备注
	ExtInfo         string `json:"ext_info"`         // 扩展信息
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetAPImethodName() string {
	return "koubei.marketing.tool.points.update"
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetAPIversion() string {
	return "1.0"
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *KoubeiMarketingToolPointsUpdateRequest) IsNeedEncrypt() bool {
	return false
}

func (this *KoubeiMarketingToolPointsUpdateRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
