package delivery

import (
	mware "github.com/arroganthooman/library-system/internal/middleware"
	userUsecase "github.com/arroganthooman/library-system/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router     *gin.Engine
	usecase    *userUsecase.Usecase
	middleware *mware.Middleware
}

func NewLibraryHandler(router *gin.Engine, usecase *userUsecase.Usecase, middleware *mware.Middleware) *Handler {
	return &Handler{
		router:     router,
		usecase:    usecase,
		middleware: middleware,
	}
}

func (h *Handler) SetEndpoint() {
	h.router.GET("/user/info", h.middleware.AuthMiddleware(), h.GetUserInfoFromAuthToken)
	h.router.POST("/user/login", h.Login)
	h.router.POST("/user/register", h.RegisterUser)
	h.router.PUT("/user/edit", h.middleware.AuthMiddleware(), h.EditUser)
}
