package zhima

import (
	"github.com/vanishs/antsdk/api"
)

type ZhimaDataFeedbackurlQueryResponse struct {
	api.AlipayResponse
	FeedbackURL string `json:"feedback_url"` // 反馈模板地址
}
