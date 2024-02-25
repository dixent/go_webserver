package repositories

import (
	"go_webserver/db"
	"go_webserver/internal/shop/models"
	"log"
)

func CreateUser(user *models.User) (int64, error) {
	rows, err := db.Connection.NamedQuery(
		"INSERT INTO users (email, password) VALUES (:email, :password) RETURNING id",
		&user,
	)
	if err != nil {
		return 0, err
	}

	var id int64
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func GetUserById(id int64) (*models.User, error) {
	user := models.User{}
	db.Connection.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	return &user, nil
}

func GetUsers() ([]*models.User, error) {
	users := []*models.User{}
	err := db.Connection.Select(&users, "SELECT * FROM users")

	if err != nil {
		log.Println("Error getting users")
		return nil, err
	}

	return users, nil
}

func GetUsersWithShops() ([]*models.User, error) {
	users := []*models.User{}
	rows, err := db.Connection.Queryx(
		`
SELECT users.id, users.email, shops.id AS shop_id, shops.name AS shop_name
FROM users JOIN shops ON users.id = shops.owner_id ORDER BY users.id
`,
	)
	var lastUser *models.User

	if err != nil {
		log.Println("Error getting users with shops")
		return nil, err
	}

	for rows.Next() {
		newUser := models.User{}
		shop := models.Shop{}
		err := rows.Scan(
			&newUser.Id, &newUser.Email, &shop.Id, &shop.Name,
		)

		if err != nil {
			log.Println("Error scan user struct")
			return nil, err
		}

		rows.StructScan(&shop)

		if lastUser == nil || lastUser.Id != newUser.Id {
			lastUser = &newUser
			lastUser.Shops = append(lastUser.Shops, shop)
			users = append(users, lastUser)
		} else {
			lastUser.Shops = append(lastUser.Shops, shop)
		}
	}

	return users, nil
}

func GetUsersWithShops2Queries() ([]*models.User, error) {
	users, err := GetUsers()

	if err != nil {
		return nil, err
	}

	ownerIds := make([]int64, len(users))

	for i, user := range users {
		ownerIds[i] = user.Id
	}

	shops, err := GetShopByIds(ownerIds)

	if err != nil {
		return nil, err
	}

	shopsByOwnerId := make(map[int64][]models.Shop)

	for _, shop := range shops {
		shopsByOwnerId[shop.OwnerId] = append(shopsByOwnerId[shop.OwnerId], shop)
	}

	for ownerId, shops := range shopsByOwnerId {
		var userIndex int

		for i, user := range users {
			if user.Id == ownerId {
				userIndex = i
				break
			}
		}

		users[userIndex].Shops = append(users[userIndex].Shops, shops...)
	}

	return users, nil
}
