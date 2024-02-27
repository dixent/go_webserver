package jwt

import (
	"testing"
)

type MockSource struct {
}

func (s *MockSource) GetJwtClaims() map[string]any {
	return map[string]any{"id": int64(1), "email": "TestEmail@gmail.com"}
}

var validator *TokenValidator

func init() {
	validator = NewTokenValidator()
}

func TestEncryptTokenSuccess(t *testing.T) {
	token, err := validator.EncryptToken(&MockSource{})

	if err != nil {
		t.Errorf(`ERROR != nil; must be nil`)
	}

	if token == "" {
		t.Errorf(`TOKEN = ""; must be not ""`)
	}
}

func TestVerifyTokenSuccess(t *testing.T) {
	token, _ := validator.EncryptToken(&MockSource{})
	err := validator.VerifyToken(token)

	if err != nil {
		t.Errorf(`ERROR is not nil; must be nil`)
	}
}

func TestVerifyTokenFailure(t *testing.T) {
	err := validator.VerifyToken("invalid token")

	if err == nil {
		t.Errorf(`ERROR is nil; must be "token is malformed: token contains an invalid number of segments"`)
	}
}
