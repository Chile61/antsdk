package api

import (
	"github.com/vanishs/antsdk/utils"
)

// IAlipayUploadRequest 上传请求接口
type IAlipayUploadRequest interface {
	IAlipayRequest
	GetFileParams() map[string]*utils.FileItem
}
