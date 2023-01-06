package delivery

import (
	"net/http"

	"github.com/arroganthooman/library-system/internal/repository"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLibraryByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, repository.Library{
		ID:       1,
		Name:     "blabla",
		Location: "bandung",
	})
}
