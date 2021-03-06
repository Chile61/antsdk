package user

import "github.com/vanishs/antsdk/api"

// AlipayUserUserinfoShareResponse alipay.user.userinfo.share(支付宝钱包用户信息共享) 返回
type AlipayUserUserinfoShareResponse struct {
	E                  api.Exception
	UserID             string `json:"alipay_user_id"`       // 用户的userId
	Avatar             string `json:"avatar"`               // 用户头像
	UserType           string `json:"user_type_value"`      // 用户类型（1/2） 1代表公司账户2代表个人账户
	UserStatus         string `json:"user_status"`          // 用户状态（Q/T/B/W）。 Q代表快速注册用户 T代表已认证用户 B代表被冻结账户 W代表已注册，未激活的账户
	IsCertified        string `json:"is_certified"`         // 是否通过实名认证。T是通过 F是没有实名认证
	Province           string `json:"province"`             // 省份名称。
	City               string `json:"city"`                 // 市名称。
	NickName           string `json:"nick_name"`            // 用户昵称
	IsStudentCertified string `json:"is_student_certified"` // 是否是学生
	Gender             string `json:"gender"`               // 性别（F：女性；M：男性）
}

// SetTags SetTags
func (resp *AlipayUserUserinfoShareResponse) SetTags() (successTag string, exceptionTag string, e *api.Exception) {
	successTag = "alipay_user_userinfo_share_response"
	exceptionTag = "error_response"
	e = &resp.E
	return
}