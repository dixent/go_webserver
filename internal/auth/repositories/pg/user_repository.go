package pg

import (
	"go_webserver/internal/auth/entities"
	"go_webserver/pkg/postgres"
	"log"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetByAuth(session *entities.Auth) entities.User {
	var user entities.User

	err := r.Get(
		&user,
		"SELECT id, email, password FROM users WHERE email = $1 AND password = $2",
		session.Email,
		session.Password,
	)

	if err != nil {
		log.Println("Failed getting user by session", err)
	}

	return user
}
