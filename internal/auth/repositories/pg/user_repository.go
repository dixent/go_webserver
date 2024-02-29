package pg

import (
	"go_webserver/internal/auth/entities"
	"go_webserver/pkg/postgres"
	"log"
)

type UserRepository struct {
	db *postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetByAuth(auth *entities.Auth) entities.User {
	var user entities.User

	err := r.db.Get(
		&user,
		"SELECT id, email, password FROM users WHERE email = $1 AND password = $2",
		auth.Email,
		auth.Password,
	)

	if err != nil {
		log.Println("Failed getting user by auth", err)
	}

	return user
}
