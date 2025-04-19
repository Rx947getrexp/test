package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/exp/rand"
)

// 生成签名
func MakeSignature(secret, timestamp, nonce, payload string) string {
	data := timestamp + nonce + payload
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// 随机字符串
func RandomNonceString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// 签名校验
func CheckSignature(secret, timestamp, nonce, payload, sig string) bool {
	data := timestamp + nonce + payload
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	expectedSig := hex.EncodeToString(mac.Sum(nil))
	return expectedSig == sig
}
