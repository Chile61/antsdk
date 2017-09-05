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

}
