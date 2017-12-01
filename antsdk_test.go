package antsdk

import (
	"fmt"
	"strings"
	"testing"

	"github.com/vanishs/antsdk/api/util"
)

func Test3(t *testing.T) {

	client := NewDefaultClient(ConstProdGateway, "appid", "private", "public", ConstSignTypeRSA, false)

	request := &util.AlipayUserInfoAuthRequest{}

	request.Scopes = []string{"auth_user", "auth_base"}
	request.State = "test"

	var response util.AlipayUserInfoAuthResponse

	err := client.Execute(request, &response)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		if response.E.IsSuccess() {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}
}

func Test4(t *testing.T) {
	a := `\\\"\/\b\f\n\r\t`
	b := `\"/` + "\b" + "\f" + "\n" + "\r" + "\t"
	// https://docs.microsoft.com/zh-cn/sql/relational-databases/json/how-for-json-escapes-special-characters-and-control-characters-sql-server
	a = strings.Replace(a, `\\`, `\`, -1)
	a = strings.Replace(a, `\"`, `"`, -1)
	a = strings.Replace(a, `\/`, `/`, -1)
	a = strings.Replace(a, `\b`, "\b", -1)
	a = strings.Replace(a, `\f`, "\f", -1)
	a = strings.Replace(a, `\n`, "\n", -1)
	a = strings.Replace(a, `\r`, "\r", -1)
	a = strings.Replace(a, `\t`, "\t", -1)
	if a != b {
		t.Error("err!\n")
	}

}
