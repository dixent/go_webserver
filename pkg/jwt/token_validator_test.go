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
		t.Errorf(`err IS NOT nil; EXPECT nil`)
	}

	if token == "" {
		t.Errorf(`token IS ""; EXPECT not ""`)
	}
}

func TestVerifyTokenSuccess(t *testing.T) {
	token, _ := validator.EncryptToken(&MockSource{})
	userId, err := validator.VerifyToken(token)

	if err != nil {
		t.Errorf(`err IS %s; EXPECT nil`, err)
	}

	if userId != 1 {
		t.Errorf(`userId IS %d; EXPECT 1`, userId)
	}
}

func TestVerifyTokenFailure(t *testing.T) {
	userId, err := validator.VerifyToken("invalid token")

	if err == nil {
		t.Errorf(`err IS nil; EXPECT "token is malformed: token contains an invalid number of segments"`)
	}

	if userId == 1 {
		t.Errorf(`userId IS 1; EXPECT 0`)
	}
}
