package services

import "go_webserver/internal/auth/entities"

type Authenticate interface {
	Authenticate(entities.Auth) string
}
