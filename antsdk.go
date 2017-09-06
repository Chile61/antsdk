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
	serverURL       string
	appID           string
	privatePKCS8B64 string
	format          string
	signType        string
	publicPKCS8B64  string
	encryptType     string
	charset         string
	encryptKey      string
	hash            crypto.Hash
}

// NewDefaultClient 创建一个默认的Client，可以构建相关struct后调用Execute执行该功能
func NewDefaultClient(serverURL, appID, privatePKCS8B64, publicPKCS8B64 string, signtype string) *Client {

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
		serverURL:       serverURL,
		appID:           appID,
		privatePKCS8B64: privatePKCS8B64,
		format:          constFormatJSON,
		signType:        signtype,
		publicPKCS8B64:  publicPKCS8B64,
		encryptType:     constEncryptTypeAES,
		charset:         constCharsetUTF8,
		hash:            hash,
	}

}

// Execute 传递相关request,response struct指针，执行相关方法并且返回结果
func (c *Client) Execute(request api.IAlipayRequest, response api.IAlipayResponse) error {
	return c.ExecuteWithAccessToken(request, response, "")
}

// ExecuteWithAccessToken 使用accessToken传递相关request,response struct指针，执行相关方法并且返回结果
func (c *Client) ExecuteWithAccessToken(request api.IAlipayRequest, response api.IAlipayResponse, accessToken string) error {
	return c.ExecuteWithAppAuthToken(request, response, accessToken, "")
}

// ExecuteWithAppAuthToken 使用appAuthToken,accessToken传递相关request,response struct指针，执行相关方法并且返回结果
func (c *Client) ExecuteWithAppAuthToken(request api.IAlipayRequest, response api.IAlipayResponse, accessToken, appAuthToken string) error {
	log.Println("#######")

	bResult, err := c.doPost(request, accessToken, appAuthToken)
	if err != nil {
		return err
	}
	log.Println("000000000")
	// 验签
	strResp := string(bResult)
	log.Println(strResp)

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
	log.Println("111111111", sign)

	///////////////////////////////////////////////////////
	s, e, ex := response.SetTags()
	log.Printf("AAAAAAAAAA [%s] [%s]\n", sign, result)
	value := gjson.Get(strResp, e)

	errjsonstr := value.String()
	log.Println("2222222", errjsonstr)
	if "" != errjsonstr {

		err = gjson.Unmarshal([]byte(errjsonstr), ex)
		if err != nil {
			log.Println("AAAAAAAAAB")
			return err
		}
		log.Println("33333333333")
		if !ex.IsSuccess() {
			return nil
		}

	}
	log.Println("4444444444")
	value = gjson.Get(strResp, s)

	okjsonstr := value.String()
	if "" == okjsonstr {
		return errors.New("json value is empty")
	}
	log.Println("5555555555")
	err = gjson.Unmarshal([]byte(okjsonstr), response)
	if err != nil {
		log.Println("BBBBBBBB")
		return err
	}
	log.Println("666666666666")

	///////////////////////////////////////////////////////

	// 验证签名
	isOk, err := utils.SyncVerifySign(sign, []byte(result), []byte(c.publicPKCS8B64), c.hash)
	if err != nil {
		return err
	}
	log.Println("!!!!!!!!")
	if !isOk {
		return errors.New("sign fail:sign error")
	}
	log.Println("@@@@@@@@@")

	return nil
}

func (c *Client) doPost(request api.IAlipayRequest, accessToken, appAuthToken string) ([]byte, error) {
	requestHolder, err := c.getRequestHolderWithSign(request, accessToken, appAuthToken)
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

	log.Println("xxx:", reqURL, data.Encode())

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

func getBizTextParams(request api.IAlipayRequest) *utils.AlipayHashMap {

	retval := utils.NewAlipayHashMap()
	retval.Put(constBizContentKey, utils.ToJSON(request))
	return retval

}

func getTextParams(request api.IAlipayRequest) *utils.AlipayHashMap {

	retval := utils.NewAlipayHashMap()

	t := reflect.TypeOf(request).Elem()
	v := reflect.ValueOf(request).Elem()

	for i := 0; i < v.NumField(); i++ {
		retval.Put(t.Field(i).Tag.Get("json"), v.Field(i).Interface())
	}

	return retval

}

func getParams(request api.IAlipayRequest) *utils.AlipayHashMap {
	if request.IsNeedBizContent() {
		return getBizTextParams(request)
	}
	return getTextParams(request)

}

func (c *Client) getRequestHolderWithSign(request api.IAlipayRequest, accessToken, appAuthToken string) (*utils.RequestParametersHolder, error) {
	requestHolder := utils.NewRequestParametersHolder()

	applicationParams := getParams(request)

	// 只有新接口和设置密钥才能支持加密
	if request.IsNeedEncrypt() {
		if !request.IsNeedBizContent() {
			return nil, errors.New("当前API不支持加密请求")
		}

		// 需要加密必须设置密钥和加密算法
		if c.encryptKey == "" || c.encryptType == "" {
			return nil, errors.New("API请求要求加密，则必须设置密钥和密钥类型：encryptKey=" + c.encryptKey + ",encryptType=" + c.encryptType)
		}

		encryptContent, err := utils.EncryptContent(applicationParams.Get(constBizContentKey), c.encryptType, c.encryptKey)
		if err != nil {
			return nil, err
		}

		applicationParams.Put(constBizContentKey, encryptContent)
	}

	protocalMustParams := utils.NewAlipayHashMap()

	protocalMustParams.Put(constAppIDKey, c.appID)
	protocalMustParams.Put(constSignTypeKey, c.signType)
	protocalMustParams.Put(constCharsetKey, c.charset)
	protocalMustParams.Put(constFormatKey, c.format)
	protocalMustParams.Put(constMethodKey, request.GetAPImethodName())
	protocalMustParams.Put(constVersionKey, constVersionVal)
	protocalMustParams.Put(constTimestampKey, time.Now().Format("2006-01-02 15:04:05"))
	protocalMustParams.Put(constAlipaySDKKey, constSDKversion)

	returnURL, ok := request.(interface{}).(*api.ReturnURL)
	if ok {
		protocalMustParams.Put(constReturnURLKey, returnURL.GetReturnURL())
	}

	notifyURL, ok := request.(interface{}).(*api.NotifyURL)
	if ok {
		protocalMustParams.Put(constNotifyURLKey, notifyURL.GetNotifyURL())
	}

	prodCode, ok := request.(interface{}).(*api.ProdCode)
	if ok {
		protocalMustParams.Put(constProdCodeKey, prodCode.GetProdCode())
	}

	terminalInfo, ok := request.(interface{}).(*api.TerminalInfo)
	if ok {
		protocalMustParams.Put(constTerminalInfoKey, terminalInfo.GetTerminalInfo())
	}

	terminalType, ok := request.(interface{}).(*api.TerminalType)
	if ok {
		protocalMustParams.Put(constTerminalTypeKey, terminalType.GetTerminalType())
	}

	if request.IsNeedEncrypt() {
		protocalMustParams.Put(constEncryptTypeKey, c.encryptType)
	}

	if "" != accessToken {
		protocalMustParams.Put(constAccessTokenKey, accessToken)
	}

	if "" != appAuthToken {
		protocalMustParams.Put(constAppAuthTokenKey, appAuthToken)
	}

	// to sign

	requestHolder.ApplicationParams = applicationParams
	requestHolder.ProtocalMustParams = protocalMustParams

	if c.signType != "" {
		signMap := utils.GetSignMap(requestHolder)
		sign, err := utils.Sign(signMap, []byte(c.privatePKCS8B64), c.hash)
		if err != nil {
			return nil, err
		}
		protocalMustParams.Put(constSignKey, sign)
	} else {
		protocalMustParams.Put(constSignKey, "")
	}

	return requestHolder, nil
}