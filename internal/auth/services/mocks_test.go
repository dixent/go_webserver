package services_test

import (
	"go_webserver/internal/auth/entities"
)

type UserRepositoryMock struct {
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (r *UserRepositoryMock) GetByAuth(session *entities.Auth) entities.User {
	if session.Email == "test" && session.Password == "test" {
		return entities.User{Id: 1, Email: "mocked_email@gmail.com", Password: ""}
	} else {
		return entities.User{}
	}
}
