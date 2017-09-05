package car

import (
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// 车辆驶出接口
// 上传车辆驶出信息，上传信息通过该接口提交到支付宝，支付宝返回对应的信息
type AlipayEcoMycarParkingExitinfoSyncRequest struct {
	api.IAlipayRequest
	TerminalType string                                             `json:"terminal_type"`
	TerminalInfo string                                             `json:"terminal_info"`
	ProdCode     string                                             `json:"prod_code"`
	NotifyURL    string                                             `json:"notify_url"`
	ReturnURL    string                                             `json:"return_url"`
	BizContent   AlipayEcoMycarParkingExitinfoSyncRequestBizContent `json:"biz_content"`
}

type AlipayEcoMycarParkingExitinfoSyncRequestBizContent struct {
	ParkingId string `json:"parking_id"` // 支付宝停车场ID，系统唯一
	CarNumber string `json:"car_number"` // 车牌号
	OutTime   string `json:"out_time"`   // 车辆离场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetAPImethodName() string {
	return "alipay.eco.mycar.parking.exitinfo.sync"
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetAPIversion() string {
	return "1.0"
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetTerminalType() string {
	return this.TerminalType
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetTerminalInfo() string {
	return this.TerminalInfo
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetNotifyURL() string {
	return this.NotifyURL
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetReturnURL() string {
	return this.ReturnURL
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetProdCode() string {
	return this.ProdCode
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) IsNeedEncrypt() bool {
	return false
}

func (this *AlipayEcoMycarParkingExitinfoSyncRequest) GetTextParams() *utils.AlipayHashMap {
	txtParams := utils.NewAlipayHashMap()
	txtParams.Put("biz_content", utils.ToJSON(this.BizContent))
	return txtParams
}
