package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var secret = []byte(viper.GetString("jwt.secret"))

type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return iJwtCustClaims, err
	}

	if !token.Valid {
		return iJwtCustClaims, errors.New("invalid token")
	}

	return iJwtCustClaims, nil
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
