package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// EncryptContent 加密内容
func EncryptContent(content, encryptType, encryptkey string) (string, error) {
	if encryptType == "AES" {
		byt, err := AesEncrypt(content, encryptkey)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(byt), nil
	}
	return "", errors.New("当前不支持该算法类型：encrypeType=" + encryptType)
}

// AesEncrypt AES 加密
func AesEncrypt(src, key string) ([]byte, error) {
	bKey := []byte(key)
	origData := []byte(src)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingPKCS5(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, bKey[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt AES 解密
func AesDecrypt(crypted []byte, key string) (string, error) {
	sKey := []byte(key)
	block, err := aes.NewCipher(sKey)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, sKey[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = unpaddingPKCS5(origData)
	// origData = ZeroUnPadding(origData)
	return string(origData), nil
}

func paddingPKCS5(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func unpaddingPKCS5(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
