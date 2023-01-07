package delivery

import (
	bookUsecase "github.com/arroganthooman/library-system/internal/book/usecase"
	mware "github.com/arroganthooman/library-system/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router     *gin.Engine
	usecase    *bookUsecase.Usecase
	middleware *mware.Middleware
}

func NewBookHandler(router *gin.Engine, usecase *bookUsecase.Usecase, middleware *mware.Middleware) *Handler {
	return &Handler{
		router:     router,
		usecase:    usecase,
		middleware: middleware,
	}
}

func (h *Handler) SetEndpoint() {
	h.router.GET("/books", h.GetAllBook) // Get All book
	h.router.GET("/book", h.GetBookByID) // Getbook by ID
	h.router.POST("/book", h.AddBook)    // Add book
	h.router.PUT("/book", h.EditBook)
	h.router.PATCH("/book", h.EditBook)
	h.router.DELETE("/book", h.DeleteBook)

	// Borrow
	h.router.POST("/book/borrow", h.middleware.AuthMiddleware(), h.BorrowBook)
	h.router.POST("/book/return", h.middleware.AuthMiddleware(), h.ReturnBook)
}
