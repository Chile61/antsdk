package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"sort"
)

// GetSignMap 获取参数map
func GetSignMap(requestHolder *RequestParametersHolder) map[string]string {
	singleMap := make(map[string]string)

	if requestHolder.ApplicationParams.Length > 0 {
		for k, v := range requestHolder.ApplicationParams.GetMap() {
			singleMap[k] = v
		}
	}

	if requestHolder.ProtocalMustParams.Length > 0 {
		for k, v := range requestHolder.ProtocalMustParams.GetMap() {
			singleMap[k] = v
		}
	}

	return singleMap
}

// GetSignStr 获取排序字符串
func GetSignStr(m map[string]string) string {
	// 对 key 进行升序排序
	sortedKeys := make([]string, 0)
	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	// 对 key=value 的键值对用 & 连接起来，略过空值
	sbSignStr := NewStringBuilder()
	for i, k := range sortedKeys {
		if m[k] != "" {
			sbSignStr.Append(k)
			sbSignStr.Append("=")
			sbSignStr.Append(m[k])
			if i != (len(sortedKeys) - 1) {
				sbSignStr.Append("&")
			}
		}
	}
	return sbSignStr.ToString()
}

// Sign 签名
func Sign(mReq map[string]string, privatePKCS8B64 []byte, hash crypto.Hash) (string, error) {

	// 获取待签名参数
	signStr := GetSignStr(mReq)

	block, err := base64.StdEncoding.DecodeString(string(privatePKCS8B64))
	if err != nil {
		return "", err
	}

	prk8, err := x509.ParsePKCS8PrivateKey(block)
	if err != nil {
		return "", err
	}

	pKey := prk8.(*rsa.PrivateKey)

	return RSASign(signStr, pKey, hash)
}

// RSASign RSA签名
func RSASign(origData string, privateKey *rsa.PrivateKey, hash crypto.Hash) (string, error) {
	h := hash.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, privateKey, hash, []byte(digest))
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)

	return string(data), nil
}

// SyncVerifySign 同步返回验签
func SyncVerifySign(body, sign string, alipayPublicKey []byte, hash crypto.Hash) (bool, error) {
	return RSAVerify(JSONUnescapeString(body), sign, alipayPublicKey, hash)
}

// AsyncVerifySign 异步返回验签
func AsyncVerifySign(body, alipayPublicKey []byte, hash crypto.Hash) (bool, error) {
	data, err := url.ParseQuery(string(body))
	if err != nil {
		return false, err
	}

	var m map[string]string
	m = make(map[string]string, 0)

	for k, v := range data {
		if k == "sign" || k == "sign_type" { //不要'sign'和'sign_type'
			continue
		}
		m[k] = v[0]
	}

	sign := data["sign"][0]

	//获取要进行计算哈希的sign string
	signStr := GetSignStr(m)

	return RSAVerify(signStr, sign, alipayPublicKey, hash)
}

// RSAVerify RSA 验证
func RSAVerify(src, sign string, publicPKCS8B64 []byte, hash crypto.Hash) (bool, error) {

	// 加载RSA的公钥

	// publicPKCS8B64[0] = 0
	block, err := base64.StdEncoding.DecodeString(string(publicPKCS8B64))
	if err != nil {
		return false, err
	}

	pub, err := x509.ParsePKIXPublicKey(block)
	if err != nil {
		return false, err
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("err publickey")
	}

	// 计算代签名字串的哈希
	t := hash.New()
	io.WriteString(t, src)
	digest := t.Sum(nil)

	// base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
	data, err := base64.StdEncoding.DecodeString(sign)

	if err != nil {
		return false, err
	}

	// 调用rsa包的VerifyPKCS1v15验证签名有效性

	err = rsa.VerifyPKCS1v15(rsaPub, hash, digest, data)
	if err != nil {
		return false, err
	}

	return true, nil
}

// ReadPemFile 读取PEM文件
func ReadPemFile(path string) []byte {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd
}