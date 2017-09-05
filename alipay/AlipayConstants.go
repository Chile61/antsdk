package alipay

const (

	// ConstProdGateway ConstProdGateway
	ConstProdGateway = "https://openapi.alipay.com/gateway.do"
	// ConstDeveGateway ConstDeveGateway
	ConstDeveGateway = "https://openapi.alipaydev.com/gateway.do"

	// ConstSignTypeRsa RSA签名算法，用于constSignTypeKey(当前默认值)
	ConstSignTypeRsa = "RSA"
	// ConstSignTypeRsa2 RSA2签名算法，用于constSignTypeKey
	ConstSignTypeRsa2 = "RSA2"

	// constCharsetUTF8 utf8字符编码格式，用于constCharsetKey(当前默认值)
	constCharsetUTF8 = "UTF-8"
	// constFormatJSON xml编码格式用于constFormatKey(当前默认值)
	constFormatJSON = "json"
	// constSDKversion SDK版本,用于constVersionKey
	constSDKversion = "alipay-sdk-go"
	// constEncryptTypeAES 加密类型 目前只支持AES
	constEncryptTypeAES = "AES"
	// constVersionVal 当前版本
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
