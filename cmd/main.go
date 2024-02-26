package main

import (
	"fmt"
	"go_webserver/config"
	"go_webserver/config/db"
	"go_webserver/internal/shop/models"
	"go_webserver/internal/shop/repositories/pg"
	"log"
	"math/rand"
)

func main() {
	config.InitEnvironment()
	db.Connection = db.InitConnection()

	getShops()

	defer db.Connection.Close()
}

func getShops() {
	repo := pg.NewShopRepository()
	if shops, err := repo.GetShops(); err != nil {
		log.Println("Error getting shops")
		panic(err)
	} else {
		for _, shop := range shops {
			log.Printf("%+v\n", shop)
		}
	}
}

func showUsers() {
	repo := pg.NewUserRepository()
	if users, err := repo.GetUsersWithShops2Queries(); err != nil {
		log.Println("Error getting users with shops")
		panic(err)
	} else {
		for _, user := range users {
			log.Printf("%+v\n", user)
		}
	}
}

func initShopsForUsers() {
	userRepo := pg.NewUserRepository()
	shopRepo := pg.NewShopRepository()
	users, err := userRepo.GetUsers()

	if err != nil {
		log.Println("Error getting users")
		panic(err)
	}

	for _, user := range users {
		shopRepo.CreateShop(user.Id, &models.Shop{Name: fmt.Sprintf("%s's shop", user.Email)})
		shopRepo.CreateShop(user.Id, &models.Shop{Name: fmt.Sprintf("%s's shop test", user.Email)})
	}
}

func createUser() {
	repo := pg.NewUserRepository()

	user := models.User{
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
