package social

import "github.com/vanishs/antsdk/api"

// AlipaySocialBaseRelationFriendsQueryResponse AlipaySocialBaseRelationFriendsQueryResponse
type AlipaySocialBaseRelationFriendsQueryResponse struct {
	E          api.Exception
	FriendList []FriendListVO `json:"friend_list"` // 好友列表数据
}

// FriendListVO FriendListVO
type FriendListVO struct {
	UserID     string `json:"user_id"`     // 用户id
	ViewName   string `json:"view_name"`   // 有可能包含emoji表情，业务方要注意编码
	HeadImg    string `json:"head_img"`    // 头像的服务地址
	RealFriend bool   `json:"real_friend"` // 是否双向好友
}

// SetTags SetTags
func (resp *AlipaySocialBaseRelationFriendsQueryResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_social_base_relation_friends_query_response"
	exceptionTag = "alipay_social_base_relation_friends_query_response"
	e = &resp.E
	return
}
