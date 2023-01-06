package delivery

import (
	"net/http"

	"github.com/arroganthooman/library-system/internal/repository"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBookByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, repository.Book{
		ID: 1,
	})
}
