package car

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 车牌查询接口
// 商户通过接口调用，获取用户车牌信息
type AlipayEcoMycarParkingVehicleQueryRequest struct {
	api.IAlipayRequest
	TerminalType string                                             `json:"terminal_type"`
	TerminalInfo string                                             `json:"terminal_info"`
	ProdCode     string                                             `json:"prod_code"`
	NotifyURL    string                                             `json:"notify_url"`
	ReturnURL    string                                             `json:"return_url"`
	BizContent   AlipayEcoMycarParkingVehicleQueryRequestBizContent `json:"biz_content"`
}

type AlipayEcoMycarParkingVehicleQueryRequestBizContent struct {
	CarId string `json:"car_id"` // 支付宝用户车辆ID，系统唯一。（该参数会在停车平台用户点击查询缴费，跳转到ISV停车缴费查询页面时，从请求中传递）
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetAPImethodName() string {
	return "alipay.eco.mycar.parking.vehicle.query"
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayEcoMycarParkingVehicleQueryRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
