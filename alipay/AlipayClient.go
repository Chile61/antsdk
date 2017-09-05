package alipay

import (
	"crypto"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/vanishs/antsdk/api"
	"github.com/vanishs/antsdk/utils"
)

// Client 是用户接口，使用NewDefaultClient可以创建一个默认的接口
type Client struct {
	serverURL       string
	appID           string
	privateKey      string
	format          string
	signType        string
	alipayPublicKey string
	encryptType     string
	charset         string
	encryptKey      string
	hash            crypto.Hash
}

// NewDefaultClient 创建一个默认的Client，可以构建相关struct后调用Execute执行该功能
func NewDefaultClient(serverURL, appID, privateKey, alipayPublicKey string, signtype string) *Client {

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
		privateKey:      privateKey,
		format:          constFormatJSON,
		signType:        signtype,
		alipayPublicKey: alipayPublicKey,
		encryptType:     constEncryptTypeAES,
		charset:         constCharsetUTF8,
		hash:            hash,
	}

}

// Execute 传递相关request,response struct指针，执行相关方法并且返回结果
func (c *Client) Execute(request, response interface{}) error {
	return c.ExecuteWithAccessToken(request, response, "")
}

// ExecuteWithAccessToken 使用accessToken传递相关request,response struct指针，执行相关方法并且返回结果
func (c *Client) ExecuteWithAccessToken(request, response interface{}, accessToken string) error {
	return c.ExecuteWithAppAuthToken(request, response, accessToken, "")
}

// ExecuteWithAppAuthToken 使用appAuthToken,accessToken传递相关request,response struct指针，执行相关方法并且返回结果
func (c *Client) ExecuteWithAppAuthToken(request, response interface{}, accessToken, appAuthToken string) error {
	bResult, err := c.doPost(request, accessToken, appAuthToken)
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

	// 验证签名
	isOk, err := utils.SyncVerifySign(sign, []byte(result), []byte(c.alipayPublicKey), c.hash)
	if err != nil {
		return err
	}

	if !isOk {
		return errors.New("验签失败:签名错误")
	}

	return json.Unmarshal([]byte(result), &response)
}

func (c *Client) doPost(request interface{}, accessToken, appAuthToken string) ([]byte, error) {
	requestHolder, err := c.getRequestHolderWithSign(reflect.ValueOf(request).Interface().(api.IAlipayRequest), accessToken, appAuthToken)
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

func getBizTextParams(request api.IAlipayRequest) *utils.AlipayHashMap {

	retval := utils.NewAlipayHashMap()
	retval.Put("biz_content", utils.ToJSON(request))
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
		sign, err := utils.Sign(signMap, []byte(c.privateKey), c.hash)
		if err != nil {
			return nil, err
		}
		protocalMustParams.Put(constSignKey, sign)
	} else {
		protocalMustParams.Put(constSignKey, "")
	}

	return requestHolder, nil
}
