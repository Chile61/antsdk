package api

// IAlipayRequest 请求接口 基本接口
type IAlipayRequest interface {
	GetAPImethodName() string
	IsNeedEncrypt() bool
	IsNeedBizContent() bool
}

// PublicTerminalType PublicTerminalType
type PublicTerminalType string

// PublicTerminalInfo PublicTerminalInfo
type PublicTerminalInfo string

// PublicNotifyURL PublicNotifyURL
type PublicNotifyURL string

// PublicReturnURL PublicReturnURL
type PublicReturnURL string

// PublicProdCode PublicProdCode
type PublicProdCode string

// PublicAuthToken PublicAuthToken
type PublicAuthToken string

// PublicAppAuthToken PublicAppAuthToken
type PublicAppAuthToken string