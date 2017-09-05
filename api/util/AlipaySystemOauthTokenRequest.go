package util

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 换取授权访问令牌
type AlipaySystemOauthTokenRequest struct {
	api.IAlipayRequest
	TerminalType string `json:"terminal_type"`
	TerminalInfo string `json:"terminal_info"`
	ProdCode     string `json:"prod_code"`
	NotifyURL    string `json:"notify_url"`
	ReturnURL    string `json:"return_url"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

func (this *AlipaySystemOauthTokenRequest) GetAPImethodName() string {
	return "alipay.system.oauth.token"
}

func (this *AlipaySystemOauthTokenRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipaySystemOauthTokenRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipaySystemOauthTokenRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipaySystemOauthTokenRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipaySystemOauthTokenRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipaySystemOauthTokenRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipaySystemOauthTokenRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipaySystemOauthTokenRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("code", this.Code)
	txtParams.Put("grant_type", this.GrantType)
	txtParams.Put("refresh_token", this.RefreshToken)
	return txtParams
}
