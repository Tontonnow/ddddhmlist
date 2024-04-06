package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var (
	Key = "secret"
)

func ParseJwtWithClaims(jwtStr string, options ...jwt.ParserOption) (jwt.Claims, error) {
	mc := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtStr, mc, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	}, options...)
	if err != nil {
		return token.Claims, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}
func MakeJwtWithClaims(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(Key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
