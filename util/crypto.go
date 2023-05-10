package util

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/ripemd160"
	"io"
	"math/big"
	"strings"
)

const addressChecksumLen = 4 //定义checksum长度为四个字节
var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

const PKCS1PrikeyType = "RSA PRIVATE KEY"
const PKCS8PrikeyType = "PRIVATE KEY"

func Ripemd160Hash(publicKey []byte) []byte {

	//将传入的公钥进行256运算，返回256位hash值
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hash := hash256.Sum(nil)

	//将上面的256位hash值进行160运算，返回160位的hash值
	ripemd160 := ripemd160.New()
	ripemd160.Write(hash)
	return ripemd160.Sum(nil) //返回Pub Key hash
}

// CheckSum 取前4个字节
func CheckSum(payload []byte) []byte {
	//这里传入的payload其实是version+Pub Key hash，对其进行两次256运算
	hash1 := sha256.Sum256(payload)

	hash2 := sha256.Sum256(hash1[:])

	return hash2[:addressChecksumLen] //返回前四个字节，为CheckSum值
}

// Base58Encode Base58编码
func Base58Encode(input []byte) []byte {
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabets[mod.Int64()])
	}
	reverseBytes(result)
	return result
}

// Base58Decode Base58解码
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(base58Alphabets, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded
}

// ReBytes ReverseBytes 翻转字节
func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func MD5(s string) string {
	m := md5.New()
	_, _ = io.WriteString(m, s)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func Guid() string {
	uid := uuid.New().String()
	return strings.ReplaceAll(uid, "-", "")
}

// GenRsaKey RSA公钥私钥产生
func GenRsaKey(bits int) (string, string, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateKeyBytes := pem.EncodeToMemory(block)
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	publicKeyBytes := pem.EncodeToMemory(block)

	return string(privateKeyBytes), string(publicKeyBytes), nil
}

// RsaEncrypt RSA加密
func RsaEncrypt(plainText []byte, publicKeyBytes []byte) ([]byte, error) {
	// pem解码
	block, _ := pem.Decode(publicKeyBytes)
	// x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	// 对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	// 返回密文
	return cipherText, nil
}

// RSA解密
// cipherText 需要解密的byte数据
// path 私钥文件路径
func RsaDecrypt(cipherText []byte, privateKeyBytes []byte) ([]byte, error) {
	// pem解码
	var priKey *rsa.PrivateKey
	block, _ := pem.Decode(privateKeyBytes)
	switch block.Type {
	case PKCS8PrikeyType:
		tmpPrikey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return []byte{}, err
		}
		priKey = tmpPrikey.(*rsa.PrivateKey)
	case PKCS1PrikeyType:
		tmpPrikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return []byte{}, err
		}
		priKey = tmpPrikey
	default:
		return []byte{}, errors.New("私钥格式错误")
	}
	// 对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherText)
	if err != nil {
		return nil, err
	}
	// 返回明文
	return plainText, nil
}

// RsaSignV2 RSA签名
func RsaSignV2(plainText []byte, privateKeyBytes []byte) (string, error) {
	// pem解码
	var priKey *rsa.PrivateKey
	block, _ := pem.Decode(privateKeyBytes)
	switch block.Type {
	case PKCS8PrikeyType:
		tmpPrikey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		priKey = tmpPrikey.(*rsa.PrivateKey)
	case PKCS1PrikeyType:
		tmpPrikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		priKey = tmpPrikey
	default:
		return "", errors.New("私钥格式错误")
	}

	// 对明文进行签名
	h := sha1.New()
	h.Write(plainText)
	digest := h.Sum(nil)
	signText, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA1, digest)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signText), nil
}

// RsaVerifyV2 RSA验签
func RsaVerifyV2(plainText []byte, publicKeyBytes []byte, signText string) error {
	sig, err := base64.StdEncoding.DecodeString(signText)
	if err != nil {
		return err
	}
	// pem解码
	block, _ := pem.Decode(publicKeyBytes)
	// x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	// 对签名进行验签
	h := sha1.New()
	h.Write(plainText)
	digest := h.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, digest, sig)
	return err
}
