package social

import "github.com/vanishs/antsdk/api"

// AlipaySocialBaseRelationFriendsQueryRequest alipay.social.base.relation.friends.query(获取好友列表信息)
type AlipaySocialBaseRelationFriendsQueryRequest struct {
	api.PublicAppAuthToken
	api.PublicAuthToken
	GetType     int64 `json:"get_type"`     // 必选 获取类型。1=获取双向好友 2=获取双向+单向好友
	IncludeSelf bool  `json:"include_self"` // 可选 好友列表中是否返回自己，true=返回 false=不返回 默认false
}

// GetAPImethodName GetAPImethodName
func (requ *AlipaySocialBaseRelationFriendsQueryRequest) GetAPImethodName() string {
	return "alipay.social.base.relation.friends.query"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipaySocialBaseRelationFriendsQueryRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipaySocialBaseRelationFriendsQueryRequest) IsNeedBizContent() bool {
	return true
}
