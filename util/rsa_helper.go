package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"sort"
	"strconv"
	"strings"
)

const (
	PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	PEM_END   = "\n-----END RSA PRIVATE KEY-----"

	PEM_PUBLIC_BEGIN = "-----BEGIN PUBLIC KEY-----\n"
	PEM_PUBLIC_END   = "\n-----END PUBLIC KEY-----"

	SIGN_IGNORE_FILED_LIST = "merchantSign"
)

/**
 * 签名
 * @param requestData	请求参数
 * @param privateKey	私钥
 * @return  string  签名串
 */
func RsaSign(requestData string, privateKey string) string {
	hash := crypto.SHA256
	shaNew := hash.New()
	shaNew.Write([]byte(requestData))
	hashed := shaNew.Sum(nil)
	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(signature)
}

/**
 * 验签
 * @param requestData	请求参数
 * @param publicKey	公钥
 * @param signature	签名串
 * @return  bool
 */
func RsaVerify(requestData string, publicKey string, signature string) bool {
	hash := crypto.SHA256
	sign, error := base64.StdEncoding.DecodeString(signature)
	if error != nil {
		return false
	}
	h := hash.New()
	h.Write([]byte(requestData))
	rsaPublicKey, _ := ParsePublicKey(publicKey)
	return rsa.VerifyPKCS1v15(rsaPublicKey, hash, h.Sum(nil), sign) == nil
}

func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateKey = FormatPrivateKey(privateKey)
	// 2、解码私钥字节，生成加密对象
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	var priKey *rsa.PrivateKey
	var priKeyTemp interface{}
	var err error
	priKeyTemp, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priKey = priKeyTemp.(*rsa.PrivateKey)
	return priKey, nil
}

func FormatPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
		privateKey = PEM_BEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEM_END) {
		privateKey = privateKey + PEM_END
	}
	return privateKey
}

func ParsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	publicKey = FormatPublicKey(publicKey)
	var rsaPublicKey *rsa.PublicKey
	var publickKeyTemp interface{}
	var err error
	block, _ := pem.Decode([]byte(publicKey))
	publickKeyTemp, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey = publickKeyTemp.(*rsa.PublicKey)
	return rsaPublicKey, nil
}

func FormatPublicKey(publicKey string) string {
	if !strings.HasPrefix(publicKey, PEM_PUBLIC_BEGIN) {
		publicKey = PEM_PUBLIC_BEGIN + publicKey
	}
	if !strings.HasSuffix(publicKey, PEM_PUBLIC_END) {
		publicKey = publicKey + PEM_PUBLIC_END
	}
	return publicKey
}

/**
 * 验证 RSA 签名
 *
 * @param publicKey
 * @param jsonStr
 * @param rawSign
 * @return
 */
func VerifyRSASign(publicKey string, jsonStr string, rawSign string) bool {
	clearTextSign := GenerateClearTextSign(jsonStr)
	return RsaVerify(clearTextSign, publicKey, rawSign)
}

/**
 * 根据json串生成明文签名
 *
 * @param jsonStr
 * @return
 */
func GenerateClearTextSign(jsonStr string) string {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &tempMap)
	if err != nil {
		//log.Error(err)
		return ""
	}
	var keys []string
	for k := range tempMap {
		if k != SIGN_IGNORE_FILED_LIST {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var text string
	for i := range keys {
		text += keys[i] + "=" + Strval(tempMap[keys[i]]) + ","
	}
	clearText := text[0 : len(text)-1]
	//log.Debug("签名明文:", clearText)
	return clearText
}

/**
 * 生成签名
 *
 * @param privateKey 商家私钥
 * @param resp       参数json格式
 * @return
 */
func GenerateSign(privateKey string, resp string) string {
	signature := RsaSign(GenerateClearTextSign(resp), privateKey)
	return signature
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
