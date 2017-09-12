package antsdk

import (
	"crypto"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

const (

	// ConstProdGateway ConstProdGateway
	ConstProdGateway = "https://openapi.alipay.com/gateway.do"
	// ConstDeveGateway ConstDeveGateway
	ConstDeveGateway = "https://openapi.alipaydev.com/gateway.do"

	// ConstSignTypeRSA RSA签名算法，用于constSignTypeKey
	ConstSignTypeRSA = "RSA"
	// ConstSignTypeRSA2 RSA2签名算法，用于constSignTypeKey(当前默认值)
	ConstSignTypeRSA2 = "RSA2"

	// constCharsetUTF8 utf8字符编码格式
	constCharsetUTF8 = "UTF-8"
	// constFormatJSON xml编码格式
	constFormatJSON = "json"
	// constSDKversion SDK版本
	constSDKversion = "alipay-sdk-go-mini"
	// constEncryptTypeAES 加密类型
	constEncryptTypeAES = "AES"
	// constVersionVal 当前API版本
	constVersionVal = "1.0"

	constSignTypeKey     = "sign_type"
	constAppIDKey        = "app_id"
	constFormatKey       = "format"
	constMethodKey       = "method"
	constTimestampKey    = "timestamp"
	constVersionKey      = "version"
	constSignKey         = "sign"
	constAlipaySDKKey    = "alipay_sdk"
	constAccessTokenKey  = "auth_token"
	constAppAuthTokenKey = "app_auth_token"
	constCharsetKey      = "charset"
	constEncryptTypeKey  = "encrypt_type"
	constBizContentKey   = "biz_content"
	constProdCodeKey     = "prod_code"
	constTerminalTypeKey = "terminal_type"
	constTerminalInfoKey = "terminal_info"
	constNotifyURLKey    = "notify_url"
	constReturnURLKey    = "return_url"

	// ConstSignAlgorithms       = "SHA1WithRSA"
	// ConstSignAlgorithmsSHA256 = "SHA256WithRSA"
	// ConstDataTimeFormat       = "yyyy-MM-dd HH:mm:ss"
	// ConstDataTimezone         = "GMT+8"
	// ConstErrorRespone         = "error_response"
	// ConstResponeSuffix        = "_response"
	// ConstResponeXMLencryptNodeName = "response_encrypted"
)

// Client 是用户接口，使用NewDefaultClient可以创建一个默认的接口
type Client struct {
	serverURL            string
	appID                string
	appPrivatePKCS8B64   string
	format               string
	signType             string
	alipayPublicPKCS8B64 string
	encryptType          string
	charset              string
	encryptKey           string
	hash                 crypto.Hash
}

// NewDefaultClient 创建一个默认的Client，可以构建相关struct后调用Execute执行该功能
func NewDefaultClient(serverURL, appID, appPrivatePKCS8B64, alipayPublicPKCS8B64 string, signtype string) *Client {

	var hash crypto.Hash

	switch signtype {
	case "RSA":
		hash = crypto.SHA1
	case "RSA2":
		hash = crypto.SHA256
	default:
		hash = crypto.SHA256
		signtype = "RSA2"
	}

	return &Client{
		serverURL:            serverURL,
		appID:                appID,
		appPrivatePKCS8B64:   appPrivatePKCS8B64,
		format:               constFormatJSON,
		signType:             signtype,
		alipayPublicPKCS8B64: alipayPublicPKCS8B64,
		encryptType:          constEncryptTypeAES,
		charset:              constCharsetUTF8,
		hash:                 hash,
	}

}

// NewGetStringClient 创建一个只能调用GetString方法的Client
func NewGetStringClient(appID, appPrivatePKCS8B64, signtype string) *Client {

	var hash crypto.Hash

	switch signtype {
	case "RSA":
		hash = crypto.SHA1
	case "RSA2":
		hash = crypto.SHA256
	default:
		hash = crypto.SHA256
		signtype = "RSA2"
	}

	return &Client{
		appID:              appID,
		appPrivatePKCS8B64: appPrivatePKCS8B64,
		format:             constFormatJSON,
		signType:           signtype,
		encryptType:        constEncryptTypeAES,
		charset:            constCharsetUTF8,
		hash:               hash,
	}

}

// GetString 获取request字符串
func (c *Client) GetString(request api.IAlipayRequest) (string, error) {
	requestHolder, err := c.getRequestHolderWithSign(request)
	if err != nil {
		return "", err
	}

	return utils.BuildQuery(requestHolder.ProtocalMustParams.GetMap()), nil
}

// Execute 传递相关request,response struct指针，执行相关方法并且返回结果
func (c *Client) Execute(request api.IAlipayRequest, response api.IAlipayResponse) error {

	bResult, err := c.doPost(request)
	if err != nil {
		return err
	}

	// 验签
	strResp := string(bResult)

	// 正则验签
	expResult := `(^\{\"[a-z|_]+\":)|(,\"sign\":\"[a-zA-Z0-9|\+|\/|\=]+\"\}$)`
	exptSign := `\"sign\":\"([a-zA-Z0-9|\+|\/|\=]+)\"`
	regResult := regexp.MustCompile(expResult)
	result := regResult.ReplaceAllString(strResp, "")
	regSign := regexp.MustCompile(exptSign)
	signMatchRes := regSign.FindStringSubmatch(strResp)
	if len(signMatchRes) < 2 {
		return errors.New("验签失败:签名丢失")
	}
	sign := signMatchRes[1]

	///////////////////////////////////////////////////////
	s, e, ex := response.SetTags()

	value := gjson.Get(strResp, e)

	errjsonstr := value.String()

	if "" != errjsonstr {

		err = gjson.Unmarshal([]byte(errjsonstr), ex)
		if err != nil {
			return err
		}
		if !ex.IsSuccess() {
			return nil
		}

	}

	value = gjson.Get(strResp, s)

	okjsonstr := value.String()
	if "" == okjsonstr {
		return errors.New("json value is empty")
	}
	err = gjson.Unmarshal([]byte(okjsonstr), response)
	if err != nil {
		return err
	}

	///////////////////////////////////////////////////////

	// 验证签名
	isOk, err := utils.SyncVerifySign(result, sign, []byte(c.alipayPublicPKCS8B64), c.hash)
	if err != nil {
		return err
	}
	if !isOk {
		return errors.New("sign fail:sign error")
	}

	return nil
}

func (c *Client) doPost(request api.IAlipayRequest) ([]byte, error) {
	requestHolder, err := c.getRequestHolderWithSign(request)
	if err != nil {
		return nil, err
	}

	reqURL := c.getRequestURL(requestHolder)

	return c.postRequest(reqURL, requestHolder.ApplicationParams.GetMap())
}

func (c *Client) postRequest(reqURL string, params map[string]string) ([]byte, error) {

	data := &url.Values{}

	for k, v := range params {
		if v != "" {
			data.Set(k, v)
		}
	}

	reqParams := ioutil.NopCloser(strings.NewReader(data.Encode()))
	var client http.Client
	req, _ := http.NewRequest("POST", reqURL, reqParams)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset="+c.charset)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) getRequestURL(requestHolder *utils.RequestParametersHolder) string {
	sbURL := utils.NewStringBuilder()
	sbURL.Append(c.serverURL)
	sysMustQuery := utils.BuildQuery(requestHolder.ProtocalMustParams.GetMap())

	sbURL.Append("?")
	sbURL.Append(sysMustQuery)

	return sbURL.ToString()
}

func (c *Client) getRequestHolderWithSign(request api.IAlipayRequest) (*utils.RequestParametersHolder, error) {

	requestHolder := utils.NewRequestParametersHolder()
	protocalMustParams := utils.NewAlipayHashMap()
	applicationParams := utils.NewAlipayHashMap()

	if request.IsNeedBizContent() {
		protocalMustParams.Put(constBizContentKey, utils.ToJSON(request))
	} else {
		t := reflect.TypeOf(request).Elem()
		v := reflect.ValueOf(request).Elem()

		for i := 0; i < v.NumField(); i++ {
			jsonkey := t.Field(i).Tag.Get("json")
			if "" != jsonkey {
				applicationParams.Put(jsonkey, v.Field(i).Interface())
			}
		}
	}

	if request.IsNeedEncrypt() {
		if !request.IsNeedBizContent() {
			panic(reflect.TypeOf(request).Elem().Name() + "have bug!")
		}

		// 需要加密必须设置密钥和加密算法
		if c.encryptKey == "" || c.encryptType == "" {
			return nil, errors.New("API请求要求加密，则必须设置密钥和密钥类型：encryptKey=" + c.encryptKey + ",encryptType=" + c.encryptType)
		}

		encryptContent, err := utils.EncryptContent(protocalMustParams.Get(constBizContentKey), c.encryptType, c.encryptKey)
		if err != nil {
			return nil, err
		}

		protocalMustParams.Put(constBizContentKey, encryptContent)
		protocalMustParams.Put(constEncryptTypeKey, c.encryptType)
	}

	protocalMustParams.Put(constAppIDKey, c.appID)
	protocalMustParams.Put(constSignTypeKey, c.signType)
	protocalMustParams.Put(constCharsetKey, c.charset)
	protocalMustParams.Put(constFormatKey, c.format)
	protocalMustParams.Put(constMethodKey, request.GetAPImethodName())
	protocalMustParams.Put(constVersionKey, constVersionVal)
	protocalMustParams.Put(constTimestampKey, time.Now().Format("2006-01-02 15:04:05"))
	protocalMustParams.Put(constAlipaySDKKey, constSDKversion)

	var requPubV reflect.Value
	requestTypeName := reflect.TypeOf(request).Elem().Name()

	requPubV = reflect.ValueOf(request).Elem().FieldByName("PublicTerminalType")
	if requPubV.IsValid() {
		if requPubV.String() == "" {
			log.Println("[W]: " + requestTypeName + " PublicTerminalType is empty")
		} else {
			protocalMustParams.Put(constTerminalTypeKey, requPubV.String())
		}

	}

	requPubV = reflect.ValueOf(request).Elem().FieldByName("PublicTerminalInfo")
	if requPubV.IsValid() {
		if requPubV.String() == "" {
			log.Println("[W]: " + requestTypeName + " PublicTerminalInfo is empty")
		} else {
			protocalMustParams.Put(constTerminalInfoKey, requPubV.String())
		}

	}

	requPubV = reflect.ValueOf(request).Elem().FieldByName("PublicNotifyURL")
	if requPubV.IsValid() {
		if requPubV.String() == "" {
			log.Println("[W]: " + requestTypeName + " PublicNotifyURL is empty")
		} else {
			protocalMustParams.Put(constNotifyURLKey, requPubV.String())
		}

	}

	requPubV = reflect.ValueOf(request).Elem().FieldByName("PublicReturnURL")
	if requPubV.IsValid() {
		if requPubV.String() == "" {
			log.Println("[W]: " + requestTypeName + " PublicReturnURL is empty")
		} else {
			protocalMustParams.Put(constReturnURLKey, requPubV.String())
		}

	}

	requPubV = reflect.ValueOf(request).Elem().FieldByName("PublicProdCode")
	if requPubV.IsValid() {
		if requPubV.String() == "" {
			log.Println("[W]: " + requestTypeName + " PublicProdCode is empty")
		} else {
			protocalMustParams.Put(constProdCodeKey, requPubV.String())
		}

	}

	requPubV = reflect.ValueOf(request).Elem().FieldByName("PublicAuthToken")
	if requPubV.IsValid() {
		if requPubV.String() == "" {
			log.Println("[W]: " + requestTypeName + " PublicAuthToken is empty")
		} else {
			protocalMustParams.Put(constAccessTokenKey, requPubV.String())
		}

	}

	requPubV = reflect.ValueOf(request).Elem().FieldByName("PublicAppAuthToken")
	if requPubV.IsValid() {
		if requPubV.String() == "" {
			log.Println("[W]: " + requestTypeName + " PublicAppAuthToken is empty")
		} else {
			protocalMustParams.Put(constAppAuthTokenKey, requPubV.String())
		}

	}

	// to sign

	requestHolder.ApplicationParams = applicationParams
	requestHolder.ProtocalMustParams = protocalMustParams

	signMap := utils.GetSignMap(requestHolder)

	sign, err := utils.Sign(signMap, []byte(c.appPrivatePKCS8B64), c.hash)
	if err != nil {
		return nil, err
	}
	protocalMustParams.Put(constSignKey, sign)

	return requestHolder, nil
}

// GetClientFastLoginStr 生成客户端快速登陆字符串
func GetClientFastLoginStr(appID, appPrivatePKCS8B64, signtype, pid string) (string, error) {

	var hash crypto.Hash

	switch signtype {
	case "RSA":
		hash = crypto.SHA1
	case "RSA2":
		hash = crypto.SHA256
	default:
		hash = crypto.SHA256
		signtype = "RSA2"
	}

	requestHolder := utils.NewRequestParametersHolder()

	requestHolder.ProtocalMustParams = utils.NewAlipayHashMap()
	requestHolder.ApplicationParams = utils.NewAlipayHashMap()

	requestHolder.ProtocalMustParams.Put("apiname", "com.alipay.account.auth")
	requestHolder.ProtocalMustParams.Put("app_name", "mc")
	requestHolder.ProtocalMustParams.Put("auth_type", "AUTHACCOUNT")
	requestHolder.ProtocalMustParams.Put("biz_type", "openservice")
	requestHolder.ProtocalMustParams.Put("method", "alipay.open.auth.sdk.code.get")
	requestHolder.ProtocalMustParams.Put("product_id", "APP_FAST_LOGIN")
	requestHolder.ProtocalMustParams.Put("scope", "kuaijie")
	requestHolder.ProtocalMustParams.Put("target_id", strconv.FormatInt(time.Now().UnixNano(), 10))

	requestHolder.ProtocalMustParams.Put("app_id", appID)
	requestHolder.ProtocalMustParams.Put("pid", pid)

	signMap := utils.GetSignMap(requestHolder)

	sign, err := utils.Sign(signMap, []byte(appPrivatePKCS8B64), hash)
	if err != nil {
		return "", errors.New("Sign fail")
	}

	requestHolder.ProtocalMustParams.Put(constSignKey, sign)
	requestHolder.ProtocalMustParams.Put(constSignTypeKey, signtype)

	sbURL := utils.NewStringBuilder()

	sysMustQuery := utils.BuildQuery(requestHolder.ProtocalMustParams.GetMap())

	sbURL.Append(sysMustQuery)

	return sbURL.ToString(), nil

}

// GetNotify 解析异步通知,小提示:body需要使用http.Request.ParseForm() and http.Request.Form.Encode()获得
func GetNotify(body string, alipayPublicKeyRSA, alipayPublicKeyRSA2 []byte) (*AlipayNotify, error) {
	var n AlipayNotify
	ok, err := utils.AsyncVerifySign(body, alipayPublicKeyRSA, alipayPublicKeyRSA2, &n)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return &n, nil

}
