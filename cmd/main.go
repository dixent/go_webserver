package main

import (
	"fmt"
	"go_webserver/db"
	"go_webserver/internal/shop/models"
	"go_webserver/internal/shop/repositories"
	"log"
	"math/rand"
)

func main() {
	db.Connection = db.InitConnection()

	if users, err := repositories.GetUsersWithShops2Queries(); err != nil {
		log.Println("Error getting users with shops")
		panic(err)
	} else {
		for _, user := range users {
			log.Printf("%+v\n", user)
		}
	}

	defer db.Connection.Close()
}

func initShopsForUsers() {
	users, err := repositories.GetUsers()

	if err != nil {
		log.Println("Error getting users")
		panic(err)
	}

	for _, user := range users {
		repositories.CreateShop(user.Id, &models.Shop{Name: fmt.Sprintf("%s's shop", user.Email)})
		repositories.CreateShop(user.Id, &models.Shop{Name: fmt.Sprintf("%s's shop test", user.Email)})
	}
}

func createUser() {
	user := models.User{
		Email:    fmt.Sprintf("test%d@gmail.com", rand.Intn(1000)),
		Password: "test_password",
	}

	userId, err := repositories.CreateUser(&user)

	if err != nil {
		log.Println("Error creating user")
		panic(err)
	}

	if user, err := repositories.GetUserById(userId); err != nil {
		log.Println("Error getting user")
		panic(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
}
