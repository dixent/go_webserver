package main

import (
	"fmt"
	"go_webserver/config"
	"go_webserver/internal/shop/entities"
	"go_webserver/internal/shop/repositories"
	"go_webserver/internal/shop/repositories/pg"
	"go_webserver/pkg/postgres"
	"log"
	"math/rand"
)

func main() {
	config.InitEnvironment()
	log.Println("I am here")
	postgres := postgres.NewPostgres()

	shopRepo := pg.NewShopRepository(postgres)
	userRepo := pg.NewUserRepository(postgres)

	createUser(userRepo)
	createUser(userRepo)
	createUser(userRepo)
	initShopsForUsers(userRepo, shopRepo)
	getShops(shopRepo)
	showUsers(userRepo)

	defer postgres.Close()
}

func getShops(repo repositories.ShopRepository) {
	if shops, err := repo.GetShops(); err != nil {
		log.Println("Error getting shops")
		panic(err)
	} else {
		for _, shop := range shops {
			log.Printf("%+v\n", shop)
		}
	}
}

func showUsers(repo repositories.UserRepository) {
	if users, err := repo.GetUsersWithShops2Queries(); err != nil {
		log.Println("Error getting users with shops")
		panic(err)
	} else {
		for _, user := range users {
			log.Printf("%+v\n", user)
		}
	}
}

func initShopsForUsers(userRepo repositories.UserRepository, shopRepo repositories.ShopRepository) {
	users, err := userRepo.GetUsers()

	if err != nil {
		log.Println("Error getting users")
		panic(err)
	}

	for _, user := range users {
		shopRepo.CreateShop(user.Id, &entities.Shop{Name: fmt.Sprintf("%s's shop", user.Email)})
		shopRepo.CreateShop(user.Id, &entities.Shop{Name: fmt.Sprintf("%s's shop test", user.Email)})
	}
}

func createUser(repo repositories.UserRepository) {
	user := entities.User{
		Email:    fmt.Sprintf("test%d@gmail.com", rand.Intn(1000)),
		Password: "test_password",
	}

	userId, err := repo.CreateUser(&user)

	if err != nil {
		log.Println("Error creating user")
		panic(err)
	}

	if user, err := repo.GetUserById(userId); err != nil {
		log.Println("Error getting user")
		panic(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
}
