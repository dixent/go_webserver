package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenValidator struct {
	SecretKey []byte
}

func NewTokenValidator() *TokenValidator {
	return &TokenValidator{SecretKey: []byte(os.Getenv("HMAC_SECRET"))}
}

type Source interface {
	GetJwtClaims() map[string]any
}

func (tv *TokenValidator) EncryptToken(source Source) (string, error) {
	JwtClaims := source.GetJwtClaims()
	JwtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims(JwtClaims),
	)

	tokenString, err := token.SignedString([]byte(tv.SecretKey))
	if err != nil {
		log.Println("Failed encrypting token", err)
		return "", err
	}

	return tokenString, nil
}

func (tv *TokenValidator) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tv.SecretKey, nil
	})

	if err != nil || !token.Valid {
		log.Println("Failed verify token", err)
		return err
	}

	return nil
}
