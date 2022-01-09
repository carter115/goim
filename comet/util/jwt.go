package util

import (
	"github.com/dgrijalva/jwt-go"
	"gmimo/comet/config"
	"time"
)

var jwtSecret = []byte(config.App.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.App.JwtExpire)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	// 三种crypto.Hash方案
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// SignedString():生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims():用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// Valid():验证基于时间的声明exp, iat, nbf
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
