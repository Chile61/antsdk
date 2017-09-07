package public

import "github.com/vanishs/antsdk/api"

// AlipayOpenPublicContactFollowBatchqueryRequest alipay.open.public.contact.follow.batchquery(查询服务窗联系人关注列表)
// 查询服务窗联系人关注列表
type AlipayOpenPublicContactFollowBatchqueryRequest struct {
	api.PublicAppAuthToken
	api.PublicAuthToken
}

// GetAPImethodName GetAPImethodName
func (requ *AlipayOpenPublicContactFollowBatchqueryRequest) GetAPImethodName() string {
	return "alipay.open.public.contact.follow.batchquery"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayOpenPublicContactFollowBatchqueryRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayOpenPublicContactFollowBatchqueryRequest) IsNeedBizContent() bool {
	return true
}