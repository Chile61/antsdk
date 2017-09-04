package finance

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 余利宝赎回
type MybankFinanceYulibaoCapitalRansomRequest struct {
	api.IAlipayRequest
	TerminalType string                                             `json:"terminal_type"`
	TerminalInfo string                                             `json:"terminal_info"`
	ProdCode     string                                             `json:"prod_code"`
	NotifyURL    string                                             `json:"notify_url"`
	ReturnURL    string                                             `json:"return_url"`
	BizContent   MybankFinanceYulibaoCapitalRansomRequestBizContent `json:"biz_content"`
}

type MybankFinanceYulibaoCapitalRansomRequestBizContent struct {
	FundCode   string `json:"fund_code"`   // 基金代码，必填。目前默认填001529，代表余利宝。
	Amount     int    `json:"amount"`      // 赎回的金额，以分为单位，必须为正整数。如amount=123456表示赎回1234.56元的余利宝份额。
	Currency   string `json:"currency"`    // 币种，CNY表示人民币，目前只支持人民币
	RansomMode string `json:"ransom_mode"` // 赎回模式，REALTIME表示实时，NOTREALTIME表示非实时赎回（T+1到账），仅支持这两种模式。实时赎回日累计金额小于等于500万，大于500万则要使用非实时赎回，选择非实时赎回且日累计金额小于等于500万则会自动转为实时。
	OutBizNo   string `json:"out_biz_no"`  // 余利宝赎回流水号，用于幂等控制。流水号必须长度在30到40位之间，且仅能由数字、字母、字符“-”和字符“_”组成。建议使用UUID，如“c39c24f1-73e5-497d-9b1f-0f585ae192c1”，或者使用自定义的数字流水号，如“201608150000000000000000000000000001”。
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetApiMethodName() string {
	return "mybank.finance.yulibao.capital.ransom"
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetApiVersion() string {
	return "1.0"
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) IsNeedEncrypt() bool {
	return false
}

func (this *MybankFinanceYulibaoCapitalRansomRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
