package api

// IAlipayResponse Response需要实现的接口
type IAlipayResponse interface {
	SetTags() (successTag string, exceptionTag string, e *Exception)
}

// Exception 异常结构 见https://docs.open.alipay.com/common/105806
type Exception struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

// IsSuccess 判断结果是否成功 见https://docs.open.alipay.com/common/105806
func (e *Exception) IsSuccess() bool {
	return e.Code == "" || e.Code == "10000"
}
