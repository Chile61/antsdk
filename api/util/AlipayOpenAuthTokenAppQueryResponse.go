package util

import (
	"github.com/vanishs/antsdk/api"
)

// AlipayOpenAuthTokenAppQueryResponse AlipayOpenAuthTokenAppQueryResponse
type AlipayOpenAuthTokenAppQueryResponse struct {
	api.AlipayResponse
	UserID      string `json:"user_id"`      // 授权商户的user_id
	AuthAppID   string `json:"auth_app_id"`  // 授权商户的appid
	ExpiresIn   int64  `json:"expires_in"`   // 应用授权令牌失效时间，单位到秒
	AuthMethods string `json:"auth_methods"` // 当前app_auth_token的授权接口列表
	AuthStart   string `json:"auth_start"`   // 授权生效时间
	AuthEnd     string `json:"auth_end"`     // 授权失效时间
	Status      string `json:"status"`       // valid：有效状态；invalid：无效状态
}
