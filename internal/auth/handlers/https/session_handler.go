package https

import (
	"go_webserver/internal/auth/entities"
	"go_webserver/internal/auth/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	*gin.RouterGroup
	authService services.Authenticate
}

func NewSessionHandler(router *gin.RouterGroup, authService services.Authenticate) *SessionHandler {
	handler := SessionHandler{router, authService}

	handler.POST("/session", handler.GetSession)
	return &handler
}

func (h SessionHandler) GetSession(c *gin.Context) {
	var auth entities.Auth

	if err := c.ShouldBindJSON(&auth); err != nil {
		log.Println("Failed sessions params", err)

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := h.authService.Authenticate(auth)

	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, token)
}
