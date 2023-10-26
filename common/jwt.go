// 传输token，需要修改jwt版本
package common

import (
	"assemblyline/project1023/model"
	"time"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5" //新版jwt
)

var jwtKey = []byte("a_secre_cret")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

// 获取token
func ReleaseToken(user model.UserTable) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(time.Second)),
			Issuer:    "assemblyline",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token获得claims
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, Claims, err
}
