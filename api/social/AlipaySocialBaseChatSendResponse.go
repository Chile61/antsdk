package social

import "github.com/vanishs/antsdk/api"

// AlipaySocialBaseChatSendResponse AlipaySocialBaseChatSendResponse
type AlipaySocialBaseChatSendResponse struct {
	E        api.Exception
	MsgIndex string `json:"msg_index"` // msgid+sessionId
}

// SetTags SetTags
func (resp *AlipaySocialBaseChatSendResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_social_base_chat_send_response"
	exceptionTag = "alipay_social_base_chat_send_response"
	e = &resp.E
	return
}
