package pg

import (
	"fmt"
	"go_webserver/internal/shop/entities"
	"go_webserver/pkg/postgres"
	"log"
	"strconv"
	"strings"
)

type ShopRepository struct {
	*postgres.Postgres
}

func NewShopRepository(db *postgres.Postgres) *ShopRepository {
	return &ShopRepository{db}
}

func (r *ShopRepository) CreateShop(ownerId int64, shop *entities.Shop) (int64, error) {
	rows, err := r.NamedQuery(
		fmt.Sprintf("INSERT INTO shops (name, owner_id) VALUES (:name, %d) RETURNING id", ownerId),
		&shop,
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

func (r *ShopRepository) GetShops() ([]entities.Shop, error) {
	shops := []entities.Shop{}
	err := r.Select(&shops, "SELECT * FROM shops")

	if err != nil {
		log.Printf("Error getting shops")
		return nil, err
	}

	return shops, nil
}

func (r *ShopRepository) GetShopByIds(ids []int64) ([]entities.Shop, error) {
	shops := []entities.Shop{}

	queryIds := make([]string, len(ids))

	for i, id := range ids {
		queryIds[i] = strconv.FormatInt(id, 10)
	}

	query := fmt.Sprintf("SELECT * FROM shops WHERE owner_id IN (%s)", strings.Join(queryIds, ","))

	err := r.Select(&shops, query)

	if err != nil {
		log.Printf("Error getting shops by ids")
		return nil, err
	}

	return shops, nil
}
