package services_test

import (
	"go_webserver/internal/auth/entities"
	"go_webserver/internal/auth/services"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
		expect   func(string)
	}{
		{
			name:     "successful authenticate",
			email:    "test",
			password: "test",
			expect: func(res string) {
				if res == "" {
					t.Fatal("token IS \"\" FOR valid Auth EXPECT jwt token")
				}
			},
		},
		{
			name:     "failure authenticate",
			email:    "invalid",
			password: "invalid",
			expect: func(res string) {
				if res != "" {
					t.Fatal("token IS NOT \"\" FOR invalid auth EXPECT empty string")
				}
			},
		},
	}

	s := services.NewAuthenticateService(NewUserRepositoryMock())

	for _, test := range testCases {
		auth := entities.Auth{Email: test.email, Password: test.password}

		token := s.Authenticate(auth)

		test.expect(token)
	}
}
