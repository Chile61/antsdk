# antsdk
内部版本:alipay-sdk-go-mini
实现最小功能 蚂蚁金服(支付宝) 开放平台 go-sdk
支付宝API文档:[传送门](https://docs.open.alipay.com/api)

## 安装
```bash
go get github.com/vanishs/antsdk
```

## 使用示例

```go
import (
  "fmt"
  "github.com/vanishs/antsdk/alipay"
  "github.com/vanishs/antsdk/api/trade"
)

func main() {
    client := alipay.NewDefaultClient(alipay.ConstProdGateway, "商户AppId", "商户密钥", "支付宝公钥", alipay.ConstSignTypeRsaRSA)

    request := &util.AlipayUserInfoAuthRequest{}

    request.Scopes = []string{"auth_user", "auth_base"}
    request.State = "test"

    var response util.AlipayUserInfoAuthResponse

    err := client.Execute(request, &response)
    if err != nil {
        // 错误处理
        fmt.Println(err)
    } else {
        if response.IsSuccess() {
          fmt.Println("调用成功")
        } else {
          fmt.Println("调用失败")
        }
    }
}
```

## API 支持情况
- [ ] [支付 API](#支付-api)
- [x] [会员 API](#会员-api)
- [ ] [店铺 API](#店铺-api)
- [ ] [营销 API](#营销-api)
- [ ] [生活号 API](#生活号-api)
- [ ] [芝麻信用 API](#芝麻信用-api)
- [x] [工具类 API](#工具类-api)
- [ ] [风险控制 API](#风险控制-api)
- [ ] [服务市场 API](#服务市场-api)
- [ ] [账务 API](#账务-api)
- [ ] [生活缴费 API](#生活缴费-api)
- [ ] [车主服务 API](#车主服务-api)
- [ ] [社区物业 API](#社区物业-api)
- [ ] [数据服务 API](#数据服务-api)
- [ ] [教育服务 API](#教育服务-api)
- [ ] [卡券 API](#卡券-api)
- [ ] [广告 API](#广告-api)
- [ ] [资金 API](#资金-api)
- [ ] [地铁购票 API](#地铁购票-api)
- [ ] [历史 API](#历史-api)
- [ ] [电子发票 API](#电子发票-api)
- [ ] [理财 API](#理财-api)

## 工具类 API

描述 | API
---|---
上报线下服务异常 | util.AlipayOfflineMarketReportErrorCreateRequest
查询某个应用授权AppAuthToken的授权信息 | util.AlipayOpenAuthTokenAppQueryRequest
换取应用授权令牌  | util.AlipayOpenAuthTokenAppRequest
换取授权访问令牌 | util.AlipaySystemOauthTokenRequest
用户登陆授权 | util.AlipayUserInfoAuthRequest

## 会员 API

描述 | API
---|---
支付宝钱包用户信息共享 | user.AlipayUserUserinfoShareRequest
支付宝会员授权信息查询接口 | user.AlipayUserInfoShareRequest

