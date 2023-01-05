package delivery

import (
	libraryUsecase "github.com/arroganthooman/library-system/internal/library/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	usecase *libraryUsecase.Usecase
}

func NewLibraryHandler(router *gin.Engine, usecase *libraryUsecase.Usecase) *Handler {
	return &Handler{
		router:  router,
		usecase: usecase,
	}
}

func (h *Handler) SetEndpoint() {

}
