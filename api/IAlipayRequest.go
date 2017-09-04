package api

import (
	"github.com/vanishs/antsdk/utils"
)

// IAlipayRequest 请求接口
type IAlipayRequest interface {
	GetApiMethodName() string
	GetApiVersion() string
	GetTerminalType() string
	GetTerminalInfo() string
	GetNotifyURL() string
	GetReturnURL() string
	GetProdCode() string
	GetTextParams() *utils.AlipayHashMap
	IsNeedEncrypt() bool
}
