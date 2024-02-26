package repositories

import "go_webserver/internal/shop/models"

type ShopRepository interface {
	CreateShop(int64, *models.Shop) (int64, error)
	GetShops(int64, *models.Shop) ([]models.Shop, error)
	GetShopByIds([]int64) ([]models.Shop, error)
}

type UserRepository interface {
	CreateUser(*models.User) (int64, error)
	GetUserById(int64) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUsersWithShops() ([]*models.User, error)
	GetUsersWithShops2Queries() ([]*models.User, error)
}
