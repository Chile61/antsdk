package alipay

import (
	"fmt"
	"testing"

	"github.com/vanishs/antsdk/api/util"
)

func Test1(t *testing.T) {
	x := util.AlipayOfflineMarketReportErrorCreateRequest{}
	v := getParams(&x)
	fmt.Println(v.GetMap())
}

func Test2(t *testing.T) {
	x := util.AlipayUserInfoAuthRequest{}
	v := getParams(&x)
	fmt.Println(v.GetMap())
}

func Test3(t *testing.T) {

	client := NewDefaultClient(ConstProdGateway, "appid", "private", "public", ConstSignTypeRsa)

	request := &util.AlipayUserInfoAuthRequest{}

	request.Scopes = []string{"auth_user", "auth_base"}
	request.State = "test"

	var response util.AlipayUserInfoAuthResponse

	err := client.Execute(request, &response)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		if response.IsSuccess() {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}
}
