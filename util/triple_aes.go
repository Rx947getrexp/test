package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"go-speed/global"
	"strings"
)

//加密
func AesEncrypt(orig string) string {
	key := "6109240608fe732b"
	iv := "a610a3285c883de2"
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	ivkey := []byte(iv)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, ivkey[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return strings.ToUpper(hex.EncodeToString(cryted))
	//return base64.StdEncoding.EncodeToString(cryted)

}

//解密
func AesDecrypt(cryted string) string {
	defer global.Recovery()
	key := "6109240608fe732b"
	iv := "a610a3285c883de2"
	// 转成字节数组
	crytedByte, err := hex.DecodeString(cryted)
	if err != nil {
		return ""
	}
	//crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	ivkey := []byte(iv)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return ""
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, ivkey[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//加密
func AesEncryptV2(orig string) string {
	key := "9859240608fe732b"
	iv := "b985a3285c883de2"
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	ivkey := []byte(iv)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, ivkey[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return strings.ToUpper(hex.EncodeToString(cryted))
	//return base64.StdEncoding.EncodeToString(cryted)

}

//解密
func AesDecryptV2(cryted string) string {
	defer global.Recovery()
	key := "9859240608fe732b"
	iv := "b985a3285c883de2"
	// 转成字节数组
	crytedByte, err := hex.DecodeString(cryted)
	if err != nil {
		return ""
	}
	//crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	ivkey := []byte(iv)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return ""
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, ivkey[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}
