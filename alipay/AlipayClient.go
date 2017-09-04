package alipay

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
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
	prodCode        string
	format          string
	signType        string
	encryptType     string
	encryptKey      string
	alipayPublicKey string
	charset         string
}

// SignType SignType
type SignType int

const (
	// RSA RSA
	RSA SignType = iota
	// RSA2 RSA2
	RSA2
)

// NewDefaultClient 创建一个默认的Client，可以构建相关struct后调用Execute执行该功能
func NewDefaultClient(serverURL, appID, privateKey, alipayPublicKey string, tp SignType) *Client {
	var signType string
	switch tp {
	case RSA:
		signType = ConstSignTypeRsa
	case RSA2:
		signType = ConstSignTypeRsa2
	default:
		signType = ConstSignTypeRsa2
	}
	utils.SetHash(signType)

	return &Client{
		serverURL:       serverURL,
		appID:           appID,
		privateKey:      privateKey,
		alipayPublicKey: alipayPublicKey,
		format:          ConstFormatJSON,
		signType:        ConstSignTypeRsa,
		encryptType:     ConstEncryptTypeAES,
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
	isOk, err := utils.SyncVerifySign(sign, []byte(result), []byte(c.alipayPublicKey))
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

	if fileReq, ok := request.(api.IAlipayUploadRequest); ok {
		return c.postFileRequest(reqURL, requestHolder.ApplicationParams.GetMap(), fileReq.GetFileParams())
	}
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
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) postFileRequest(reqURL string, params map[string]string, fileParams map[string]*utils.FileItem) ([]byte, error) {
	if fileParams == nil || len(fileParams) == 0 {
		return c.postRequest(reqURL, params)
	}

	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)

	if params != nil && len(params) > 0 {
		for k, v := range params {
			w.WriteField(k, v)
		}
	}

	for k, v := range fileParams {
		fw, err := w.CreateFormFile(k, v.FileName)
		if err != nil {
			return nil, err
		}
		fw.Write(v.Content)
	}
	w.Close()

	req, err := http.NewRequest("POST", reqURL, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) getRequestURL(requestHolder *utils.RequestParametersHolder) string {
	sbURL := utils.NewStringBuilder()
	sbURL.Append(c.serverURL)
	sysMustQuery := utils.BuildQuery(requestHolder.ProtocalMustParams.GetMap())
	sysOptQuery := utils.BuildQuery(requestHolder.ProtocalOptParams.GetMap())

	sbURL.Append("?")
	sbURL.Append(sysMustQuery)
	if sysOptQuery != "" {
		sbURL.Append("&")
		sbURL.Append(sysOptQuery)
	}

	return sbURL.ToString()
}

func (c *Client) getRequestHolderWithSign(request api.IAlipayRequest, accessToken, appAuthToken string) (*utils.RequestParametersHolder, error) {
	requestHolder := utils.NewRequestParametersHolder()
	appParams := request.GetTextParams()

	// 只有新接口和设置密钥才能支持加密
	if request.IsNeedEncrypt() {
		if appParams.Get(constBizContentKey) == "" {
			return nil, errors.New("当前API不支持加密请求")
		}
		// 需要加密必须设置密钥和加密算法
		if c.encryptKey == "" || c.encryptType == "" {
			return nil, errors.New("API请求要求加密，则必须设置密钥和密钥类型：encryptKey=" + c.encryptKey + ",encryptType=" + c.encryptType)
		}

		encryptContent, err := utils.EncryptContent(appParams.Get(constBizContentKey), c.encryptType, c.encryptKey)
		if err != nil {
			return nil, err
		}

		appParams.Put(constBizContentKey, encryptContent)
	}

	if appAuthToken != "" {
		appParams.Put(constAppAuthTokenKey, appAuthToken)
	}

	requestHolder.ApplicationParams = appParams

	if c.charset == "" {
		c.charset = ConstCharsetUTF8
	}

	protocalMustParams := utils.NewAlipayHashMap()
	protocalMustParams.Put(constMethodKey, request.GetApiMethodName())
	protocalMustParams.Put(constVersionKey, request.GetApiVersion())
	protocalMustParams.Put(constAppIDKey, c.appID)
	protocalMustParams.Put(constSignTypeKey, c.signType)
	protocalMustParams.Put(constTerminalTypeKey, request.GetTerminalType())
	protocalMustParams.Put(constTerminalInfoKey, request.GetTerminalInfo())
	protocalMustParams.Put(constNotifyURLKey, request.GetNotifyURL())
	protocalMustParams.Put(constReturnURLKey, request.GetReturnURL())
	protocalMustParams.Put(constCharsetKey, c.charset)

	if request.IsNeedEncrypt() {
		protocalMustParams.Put(constEncryptTypeKey, c.encryptType)
	}
	protocalMustParams.Put(constTimestampKey, time.Now().Format("2006-01-02 15:04:05"))
	requestHolder.ProtocalMustParams = protocalMustParams

	protocalOptParams := utils.NewAlipayHashMap()
	protocalOptParams.Put(constFormatKey, c.format)
	protocalOptParams.Put(constAccessTokenKey, accessToken)
	protocalOptParams.Put(constAlipaySDKKey, ConstSDKversion)
	protocalOptParams.Put(constProdCodeKey, request.GetProdCode())
	requestHolder.ProtocalOptParams = protocalOptParams

	if c.signType != "" {
		signMap := utils.GetSignMap(requestHolder)
		sign, err := utils.Sign(signMap, []byte(c.privateKey))
		if err != nil {
			return nil, err
		}
		protocalMustParams.Put(constSignKey, sign)
	} else {
		protocalMustParams.Put(constSignKey, "")
	}

	return requestHolder, nil
}
