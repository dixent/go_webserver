package jwt

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenValidator struct {
	SecretKey []byte
}

type Claims struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	jwt.MapClaims
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

func (tv *TokenValidator) VerifyToken(tokenString string) (int64, error) {
	var userId int64

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return tv.SecretKey, nil
	})

	if err != nil || !token.Valid {
		log.Println("Failed verify token", err)
		return userId, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return userId, errors.New("failed parsing claims")
	}

	return claims.Id, nil
}
