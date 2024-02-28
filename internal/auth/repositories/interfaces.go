package repositories

import (
	"go_webserver/internal/auth/entities"
)

type UserRepository interface {
	GetByAuth(*entities.Auth) entities.User
}
