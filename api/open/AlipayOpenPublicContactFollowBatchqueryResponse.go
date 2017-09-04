package open

import (
	"github.com/vanishs/antsdk/api"
)

type AlipayOpenPublicContactFollowBatchqueryResponse struct {
	api.AlipayResponse
	ContactFollowList []ContactFollower `json:"contact_follow_list"` // 联系人关注者列表
}
