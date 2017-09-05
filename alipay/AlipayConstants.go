package alipay

const (

	// ConstSignTypeRsa RSA签名算法，用于constSignTypeKey(当前默认值)
	ConstSignTypeRsa = "RSA"

	// ConstSignTypeRsa2 RSA2签名算法，用于constSignTypeKey
	ConstSignTypeRsa2 = "RSA2"

	// ConstCharsetUTF8 utf8字符编码格式，用于constCharsetKey(当前默认值)
	ConstCharsetUTF8 = "UTF-8"
	// ConstCharsetGBK utf8字符编码格式，用于constCharsetKey
	ConstCharsetGBK = "GBK"

	// ConstFormatJSON xml编码格式用于constFormatKey(当前默认值)
	ConstFormatJSON = "json"
	// ConstFormatXML xml编码格式用于constFormatKey(当前不支持)
	// ConstFormatXML = "xml"

	// ConstSDKversion SDK版本,用于constVersionKey
	ConstSDKversion = "alipay-sdk-go"

	// ConstEncryptTypeAES 加密类型 目前只支持AES
	ConstEncryptTypeAES = "AES"

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
	constTerminalTypeKey = "terminal_type"
	constTerminalInfoKey = "terminal_info"
	constCharsetKey      = "charset"
	constNotifyURLKey    = "notify_url"
	constReturnURLKey    = "return_url"
	constEncryptTypeKey  = "encrypt_type"
	constBizContentKey   = "biz_content"
	constProdCodeKey     = "prod_code"

	// ConstSignAlgorithms       = "SHA1WithRSA"
	// ConstSignAlgorithmsSHA256 = "SHA256WithRSA"
	// ConstDataTimeFormat       = "yyyy-MM-dd HH:mm:ss"
	// ConstDataTimezone         = "GMT+8"
	// ConstErrorRespone         = "error_response"
	// ConstResponeSuffix        = "_response"
	// ConstResponeXMLencryptNodeName = "response_encrypted"
)
