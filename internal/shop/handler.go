package shop

import (
	"go_webserver/internal/shop/handlers"
	"log"
	"net/http"
)

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("%s %v\n", r.Method, r.URL.Path)

	if path == "/api/shops" || path == "/api/shops/" {
		handlers.ShopHandler{}.ServeHTTP(w, r)
	}
}
