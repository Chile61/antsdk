package bill

import (
	"github.com/vanishs/antsdk/api"
)

type AlipayDataDataserviceBillDownloadurlQueryResponse struct {
	api.AlipayResponse
	BillDownloadURL string `json:"bill_download_url"` // 账单下载地址链接，获取连接后30秒后未下载，链接地址失效。
}
