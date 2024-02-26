package handlers

import (
	"encoding/json"
	"go_webserver/internal/shop/entities"
	"go_webserver/internal/shop/repositories/pg"
	"go_webserver/pkg/postgres"
	"log"
	"net/http"
)

type ShopHandler struct {
}

func (h ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/api/shops" || path == "/api/shops/" {
		switch r.Method {
		case "GET":
			indexAction(w, r)
		case "POST":
			createAction(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func indexAction(w http.ResponseWriter, _ *http.Request) {
	postgres := postgres.NewPostgres()

	repo := pg.NewShopRepository(postgres)
	if shops, err := repo.GetShops(); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		log.Println("Shops:", shops)
		json.NewEncoder(w).Encode(shops)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	defer postgres.Close()
}

func createAction(w http.ResponseWriter, r *http.Request) {
	postgres := postgres.NewPostgres()
	var shop entities.Shop
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&shop); err != nil {
		log.Println("Error getting body for createShop\n", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	repo := pg.NewShopRepository(postgres)
	if shopId, err := repo.CreateShop(shop.OwnerId, &shop); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(shopId)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	defer postgres.Close()
}
