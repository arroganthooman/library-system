package delivery

import (
	bookUsecase "github.com/arroganthooman/library-system/internal/book/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	usecase *bookUsecase.Usecase
}

func NewBookHandler(router *gin.Engine, usecase *bookUsecase.Usecase) *Handler {
	return &Handler{
		router:  router,
		usecase: usecase,
	}
}

func (h *Handler) SetEndpoint() {

}
