package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-speed/global"
	"time"
)

const (
	CommonUserType   = 1
	MerchantUserType = 2
	AdminUserType    = 3
)

var (
	jwtSecretKeyCommon   = "2079e25aab8011edbbf100ff0d5e4fc2"
	jwtSecretKeyMerchant = "2b4ab6d3ab8011ed855100ff0d5e4fc2"
	jwtSecretKeyAdmin    = "39adee1dab8011edb5d900ff0d5e4fc2"
)

// CustomClaims 自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func GenerateTokenByUser(userId, userType int64) string {
	switch userType {
	case CommonUserType:
		return GenerateToken(userId, jwtSecretKeyCommon)
	case MerchantUserType:
		return GenerateToken(userId, jwtSecretKeyMerchant)
	case AdminUserType:
		return GenerateToken(userId, jwtSecretKeyAdmin)
	default:
		return GenerateToken(userId, jwtSecretKeyCommon)
	}
}

func ParseTokenByUser(tokenString string, userType int64) (*CustomClaims, error) {
	switch userType {
	case CommonUserType:
		return ParseToken(tokenString, jwtSecretKeyCommon)
	case MerchantUserType:
		return ParseToken(tokenString, jwtSecretKeyMerchant)
	case AdminUserType:
		return ParseToken(tokenString, jwtSecretKeyAdmin)
	default:
		return ParseToken(tokenString, jwtSecretKeyCommon)
	}
}

func GenerateToken(userId int64, jwtSecretKey string) string {
	//生成token
	maxAge := 60 * 60 * 24 * 30
	customClaims := &CustomClaims{
		UserId: userId, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		global.Logger.Err(err).Msg("生成token错误")
		return ""
	}
	return tokenString

}

// ParseToken 解析token
func ParseToken(tokenString string, jwtSecretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("unknown token")
}
