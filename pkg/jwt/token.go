package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}
)

func BuildTokens(opt TokenOptions) (string, error) {
	now := time.Now().Add(-time.Minute).Unix()
	accessToken, err := genToken(now, opt.AccessSecret, opt.Fields, opt.AccessExpire)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
