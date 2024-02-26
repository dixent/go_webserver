package pg

import (
	"go_webserver/internal/shop/entities"
	"go_webserver/pkg/postgres"
	"log"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *entities.User) (int64, error) {
	rows, err := r.NamedQuery(
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

func (r *UserRepository) GetUserById(id int64) (*entities.User, error) {
	user := entities.User{}
	r.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	return &user, nil
}

func (r *UserRepository) GetUsers() ([]*entities.User, error) {
	users := []*entities.User{}
	err := r.Select(&users, "SELECT * FROM users")

	if err != nil {
		log.Println("Error getting users")
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUsersWithShops() ([]*entities.User, error) {
	users := []*entities.User{}
	rows, err := r.Queryx(
		`
SELECT users.id, users.email, shops.id AS shop_id, shops.name AS shop_name
FROM users JOIN shops ON users.id = shops.owner_id ORDER BY users.id
`,
	)
	var lastUser *entities.User

	if err != nil {
		log.Println("Error getting users with shops")
		return nil, err
	}

	for rows.Next() {
		newUser := entities.User{}
		shop := entities.Shop{}
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

func (r *UserRepository) GetUsersWithShops2Queries() ([]*entities.User, error) {
	users, err := r.GetUsers()

	if err != nil {
		return nil, err
	}

	ownerIds := make([]int64, len(users))

	for i, user := range users {
		ownerIds[i] = user.Id
	}

	shopRepo := NewShopRepository(r.Postgres)
	shops, err := shopRepo.GetShopByIds(ownerIds)

	if err != nil {
		return nil, err
	}

	shopsByOwnerId := make(map[int64][]entities.Shop)

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
