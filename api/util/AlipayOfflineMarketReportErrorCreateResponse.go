package util

import "../../api"

// AlipayOfflineMarketReportErrorCreateResponse AlipayOfflineMarketReportErrorCreateResponse
type AlipayOfflineMarketReportErrorCreateResponse struct {
	E api.Exception
}

// SetTags SetTags
func (resp *AlipayOfflineMarketReportErrorCreateResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_offline_market_reporterror_create_response"
	exceptionTag = "alipay_offline_market_reporterror_create_response"
	e = &resp.E
	return
}
