package delivery

import (
	userUsecase "github.com/arroganthooman/library-system/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	usecase *userUsecase.Usecase
}

func NewLibraryHandler(router *gin.Engine, usecase *userUsecase.Usecase) *Handler {
	return &Handler{
		router:  router,
		usecase: usecase,
	}
}

func (h *Handler) SetEndpoint() {
	h.router.POST("/login", h.Login)
	h.router.POST("/register", h.RegisterUser)
}
