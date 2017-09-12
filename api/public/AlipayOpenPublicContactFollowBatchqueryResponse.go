package public

import "../../api"

// AlipayOpenPublicContactFollowBatchqueryResponse AlipayOpenPublicContactFollowBatchqueryResponse
type AlipayOpenPublicContactFollowBatchqueryResponse struct {
	E                 api.Exception
	ContactFollowList []ContactFollower `json:"contact_follow_list"` // 联系人关注者列表
}

// ContactFollower ContactFollower
type ContactFollower struct {
	Avatar         string `json:"avatar"`           // 支付宝头像
	DefaultAvatar  bool   `json:"default_avatar"`   // 默认头像
	EachRecordFlag bool   `json:"each_record_flag"` //
	UserID         string `json:"user_id"`          // 用户id
}

// SetTags SetTags
func (resp *AlipayOpenPublicContactFollowBatchqueryResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_open_public_contact_follow_batchquery_response"
	exceptionTag = "alipay_open_public_contact_follow_batchquery_response"
	e = &resp.E
	return
}
