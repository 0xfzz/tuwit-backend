package token

import (
	"errors"
	"log"

	"github.com/0xfzz/tuwit-backend/config"
	"github.com/golang-jwt/jwt/v5"
)

func Sign(payload jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)

	tokenStr, err := token.SignedString([]byte(config.Get().JWT.Secret))

	if err != nil {
		log.Fatalf("Can't sign token, error : %s", err)
	}
	return tokenStr
}

func Decode(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad signing method")
		}
		return []byte(config.Get().JWT.Secret), nil
	})
	return t, err
}
