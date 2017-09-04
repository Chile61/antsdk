package utils

// RequestParametersHolder RequestParametersHolder
type RequestParametersHolder struct {
	ProtocalMustParams *AlipayHashMap
	ProtocalOptParams  *AlipayHashMap
	ApplicationParams  *AlipayHashMap
}

// NewRequestParametersHolder NewRequestParametersHolder
func NewRequestParametersHolder() *RequestParametersHolder {
	return &RequestParametersHolder{}
}
