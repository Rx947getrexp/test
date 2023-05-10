package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

type GoogleAuth struct {
}

func NewGoogleAuth() *GoogleAuth {
	return &GoogleAuth{}
}

//GoogleAuthenticator2FaSha1 只实现google authenticator sha1
type GoogleAuthenticator2FaSha1 struct {
	Base32NoPaddingEncodedSecret string //The base32NoPaddingEncodedSecret parameter is an arbitrary key value encoded in Base32 according to RFC 3548. The padding specified in RFC 3548 section 2.2 is not required and should be omitted.
	ExpireSecond                 uint64 //更新周期单位秒
	Digits                       int    //数字数量
}

//Totp 计算Time-based One-time Password 数字
func (m *GoogleAuthenticator2FaSha1) Totp() (code string, err error) {
	count := uint64(time.Now().Unix()) / m.ExpireSecond
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(m.Base32NoPaddingEncodedSecret)
	if err != nil {
		return "", errors.New("https://github.com/google/google-authenticator/wiki/Key-Uri-Format,REQUIRED: The base32NoPaddingEncodedSecret parameter is an arbitrary key value encoded in Base32 according to RFC 3548. The padding specified in RFC 3548 section 2.2 is not required and should be omitted.")
	}
	codeInt := hotp(key, count, m.Digits)
	intFormat := fmt.Sprintf("%%0%dd", m.Digits) //数字长度补零
	return fmt.Sprintf(intFormat, codeInt), nil
}

func hotp(key []byte, counter uint64, digits int) int {
	//RFC 6238
	//只支持sha1
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	//取sha1的最后4byte
	//0x7FFFFFFF 是long int的最大值
	//math.MaxUint32 == 2^32-1
	//& 0x7FFFFFFF == 2^31  Set the first bit of truncatedHash to zero  //remove the most significant bit
	// len(sum)-1]&0x0F 最后 像登陆 (bytes.len-4)
	//取sha1 bytes的最后4byte 转换成 uint32
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	//取十进制的余数
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}

func VerifyAuthenticator(secret, code string) bool {
	//global.Logger.Err(nil).Msgf("secret:%s,code:%s", secret, code)
	g := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: secret,
		ExpireSecond:                 30,
		Digits:                       6,
	}
	totp, err := g.Totp()
	if err != nil {
		return false
	}
	return totp == code
}

// 获取秘钥
func (this *GoogleAuth) GetSecret() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, this.un())
	return strings.ToUpper(this.base32encode(this.hmacSha1(buf.Bytes(), nil)))
}

func (this *GoogleAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (this *GoogleAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}
func (this *GoogleAuth) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

// 用于谷歌验证码加密解密
var BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

// 加密
func Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(BASE64Table)
	return coder.EncodeToString(content)
}

// 解密
func Decode(data string) string {
	coder := base64.NewEncoding(BASE64Table)
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}

//获取ip
func GetRequestIp(c *gin.Context) string {
	reqIp := c.ClientIP()
	if reqIp == "::1" {
		reqIp = "127.0.0.1"
	}
	return reqIp
}
