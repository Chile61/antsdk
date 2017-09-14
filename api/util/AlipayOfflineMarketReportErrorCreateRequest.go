package util

import "github.com/vanishs/antsdk/api"

// AlipayOfflineMarketReportErrorCreateRequest alipay.offline.market.reporterror.create(上报线下服务异常)
// 1、场景1：口碑会通过口碑详情直接跳转到ISV，同时会带桌码，如果桌码不能识别，调用接口上报，口碑会进行报警。
type AlipayOfflineMarketReportErrorCreateRequest struct {
	api.PublicAppAuthToken
	ShopID     string            `json:"shop_id"`     // 可选 口碑门店ID
	MerchantID string            `json:"merchant_id"` // 可选	商户ID
	UserID     string            `json:"user_id"`     // 可选	用户的ID
	ErrTime    int64             `json:"err_time"`    // 必选 发生错误的时候，当前系统的毫秒数，系统会把当前时间构建成Date对象保存为错误发生时间
	Type       string            `json:"type"`        // 必选 上传类型，通过类型来区分不同错误： value=tableNum 代表扫码点菜
	Feature    map[string]string `json:"feature"`     // 可选 如果：type是tableNum 请设置table_num字段作为桌码
}

// GetAPImethodName GetAPImethodName
func (requ *AlipayOfflineMarketReportErrorCreateRequest) GetAPImethodName() string {
	return "alipay.offline.market.reporterror.create"
}

// IsNeedEncrypt IsNeedEncrypt
func (requ *AlipayOfflineMarketReportErrorCreateRequest) IsNeedEncrypt() bool {
	return false
}

// IsNeedBizContent IsNeedBizContent
func (requ *AlipayOfflineMarketReportErrorCreateRequest) IsNeedBizContent() bool {
	return true
}
