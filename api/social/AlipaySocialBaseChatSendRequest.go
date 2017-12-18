package social

import "github.com/vanishs/antsdk/api"

// AlipaySocialBaseChatSendRequest alipay.social.base.chat.send(获取好友列表信息)
type AlipaySocialBaseChatSendRequest struct {
	api.PublicAppAuthToken
	api.PublicAuthToken
	ClientMsgID      string `json:"client_msg_id"`     // 必选 客户端的消息id
	ReceiverID       string `json:"receiver_id"`       // 必选 如果是个人消息，是接收消息者的userid，如果是群消息，是群的id，必填
	ReceiverUsertype string `json:"receiver_usertype"` // 必选 接受者的用户类型，支付宝1，群组2，讨论组3
	TemplateType     string `json:"template_type"`     // 必选 消息模板的类型，分享SHARE，文本TEXT，图片IMAGE
	TemplateData     string `json:"template_data"`     // 必选 消息体的内容，形式为json字符串，必填 分享模板{"title":"支付宝聊天","desc":"支付宝聊天","image":"图片地址","thumb":"缩略图地址"} 文本模板 {"m":"文本消息"}
	BizMemo          string `json:"biz_memo"`          // 必选 消息简短描述，显示在会话列表上 [链接]美副总统：美国已为迎来女总统做好准备-新闻频道-手机搜狐
	Link             string `json:"link"`              // 可选 点击消息card跳转的地址 http://www.taobao.com
}

// GetAPImethodName GetAPImethodName
func (requ *AlipaySocialBaseChatSendRequest) GetAPImethodName() string {
	return "alipay.social.base.chat.send"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipaySocialBaseChatSendRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipaySocialBaseChatSendRequest) IsNeedBizContent() bool {
	return true
}
