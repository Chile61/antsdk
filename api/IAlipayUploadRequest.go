package api

import (
  "github.com/vanishs/antsdk/utils"
)

type IAlipayUploadRequest interface {
  IAlipayRequest
  GetFileParams() map[string]*utils.FileItem
}
