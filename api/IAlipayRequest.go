package api

// IAlipayRequest 请求接口 基本接口
type IAlipayRequest interface {
	GetAPImethodName() string
	IsNeedEncrypt() bool
	IsNeedBizContent() bool
}

// ITerminalType ITerminalType
type ITerminalType interface {
	GetTerminalType() string
}

// ITerminalInfo ITerminalInfo
type ITerminalInfo interface {
	GetTerminalInfo() string
}

// INotifyURL INotifyURL
type INotifyURL interface {
	GetNotifyURL() string
}

// IReturnURL IReturnURL
type IReturnURL interface {
	GetReturnURL() string
}

// IProdCode IProdCode
type IProdCode interface {
	GetProdCode() string
}

// TerminalType TerminalType
type TerminalType struct {
	PublicTerminalType string
}

// GetTerminalType GetTerminalType
func (v *TerminalType) GetTerminalType() string {
	return v.PublicTerminalType
}

// TerminalInfo TerminalInfo
type TerminalInfo struct {
	PublicTerminalInfo string
}

// GetTerminalInfo GetTerminalInfo
func (v *TerminalInfo) GetTerminalInfo() string {
	return v.PublicTerminalInfo
}

// NotifyURL NotifyURL
type NotifyURL struct {
	PublicNotifyURL string
}

// GetNotifyURL GetNotifyURL
func (v *NotifyURL) GetNotifyURL() string {
	return v.PublicNotifyURL
}

// ReturnURL ReturnURL
type ReturnURL struct {
	PublicReturnURL string
}

// GetReturnURL GetReturnURL
func (v *ReturnURL) GetReturnURL() string {
	return v.PublicReturnURL
}

// ProdCode ProdCode
type ProdCode struct {
	PublicProdCode string
}

// GetProdCode GetProdCode
func (v *ProdCode) GetProdCode() string {
	return v.PublicProdCode
}
