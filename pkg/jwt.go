package pkg

import (
	"github.com/cicdi-go/jwt"
	"time"
)

func GenerateJwt(secret string, data map[string]interface{}, expired time.Time) (encoded string, err error) {
	algorithm := jwt.HmacSha256(secret)
	claims := jwt.NewClaim()
	claims.SetTime("exp", expired)
	for key, value := range data {
		claims.Set(key, value)
	}
	encoded, err = algorithm.Encode(claims)
	return
}

func VerifyJwt(secret string, encoded string) (claims *jwt.Claims, err error) {
	algorithm := jwt.HmacSha256(secret)
	claims, err = algorithm.DecodeAndValidate(encoded)
	return
}
