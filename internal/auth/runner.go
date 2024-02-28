package auth

import (
	"go_webserver/internal/auth/handlers/https"
	"go_webserver/internal/auth/repositories/pg"
	"go_webserver/internal/auth/services"
	"go_webserver/pkg/postgres"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Runner struct {
	mux      *http.ServeMux
	postgres *postgres.Postgres
}

func NewRunner(mux *http.ServeMux, postgres *postgres.Postgres) *Runner {
	return &Runner{mux: mux, postgres: postgres}
}

func (r *Runner) Run() {
	r.initHttpServe()
}

func (r *Runner) initHttpServe() {
	engine := gin.New()
	h := engine.Group("/api/auth")
	r.initSession(h)
	r.mux.Handle("/api/auth/", engine)
	r.mux.Handle("/api/auth", engine)
}

func (r *Runner) initSession(h *gin.RouterGroup) {
	repo := pg.NewUserRepository(r.postgres)
	service := services.NewAuthenticateService(repo)

	https.NewSessionHandler(h, service)
}
