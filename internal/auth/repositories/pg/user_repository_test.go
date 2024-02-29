package pg_test

import (
	"go_webserver/config"
	"go_webserver/internal/auth/entities"
	"go_webserver/internal/auth/repositories/pg"
	"go_webserver/pkg/postgres"
	"os"
	"testing"
)

var db *postgres.Postgres

func init() {
	os.Setenv("ENV", "test")
	config.InitEnvironment()
	db = postgres.NewPostgres()
}

func before(t *testing.T) {
	_, err := db.NamedExec(
		"INSERT INTO users (email, password) VALUES (:email, :password)",
		&entities.User{Email: "test@gmail.com", Password: "password_test"},
	)
	if err != nil {
		t.Fatal(err)
	}
}

func after(t *testing.T) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetByAuth(t *testing.T) {
	before(t)

	defer db.Close()
	defer after(t)

	testCases := []struct {
		name     string
		email    string
		password string
		expect   func(entities.User)
	}{
		{
			name:     "found by auth",
			email:    "test@gmail.com",
			password: "password_test",
			expect: func(res entities.User) {
				if res.Id == 0 {
					t.Fatal("user Id IS 0 FOR valid Auth EXPECT NOT 0")
				}
			},
		},
		{
			name:     "didn't find by auth",
			email:    "invalid",
			password: "invalid",
			expect: func(res entities.User) {
				if res.Id != 0 {
					t.Fatal("user Id IS NOT 0 FOR invalid Auth EXPECT 0")
				}
			},
		},
	}

	repo := pg.NewUserRepository(db)

	for _, test := range testCases {
		auth := entities.Auth{Email: test.email, Password: test.password}
		user := repo.GetByAuth(&auth)
		test.expect(user)
	}
}
