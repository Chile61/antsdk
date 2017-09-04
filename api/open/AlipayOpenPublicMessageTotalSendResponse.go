package open

import (
	"github.com/vanishs/antsdk/api"
)

type AlipayOpenPublicMessageTotalSendResponse struct {
	api.AlipayResponse
	MessageId string `json:"message_id"` // 消息ID
}
