package api

// AlipayResponse 支付宝返回结果结构
type AlipayResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
	Sign    string `json:"sign"`
}

// IsSuccess 判断结果是否成功
func (resp *AlipayResponse) IsSuccess() bool {
	return resp.SubCode == ""
}
