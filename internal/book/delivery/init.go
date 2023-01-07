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
	h.router.GET("/books", h.GetAllBook) // Get All book
	h.router.GET("/book", h.GetBookByID) // Getbook by ID
	h.router.POST("/book", h.AddBook)    // Add book
	h.router.PUT("/book", h.EditBook)
	h.router.DELETE("/book", h.DeleteBook)

	// Borrow
	h.router.POST("/borrow", h.BorrowBook)
}
