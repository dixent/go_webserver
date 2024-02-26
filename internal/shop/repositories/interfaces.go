package repositories

import "go_webserver/internal/shop/entities"

type ShopRepository interface {
	CreateShop(int64, *entities.Shop) (int64, error)
	GetShops() ([]entities.Shop, error)
	GetShopByIds([]int64) ([]entities.Shop, error)
}

type UserRepository interface {
	CreateUser(*entities.User) (int64, error)
	GetUserById(int64) (*entities.User, error)
	GetUsers() ([]*entities.User, error)
	GetUsersWithShops() ([]*entities.User, error)
	GetUsersWithShops2Queries() ([]*entities.User, error)
}
