package delivery

import (
	"net/http"
	"strconv"

	"github.com/arroganthooman/library-system/internal/repository"
	"github.com/arroganthooman/library-system/presentation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddBook(ctx *gin.Context) {
	req := presentation.AddBookRequestBody{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Given Request Data is not valid",
			RawErrorMessage: err.Error(),
		})
		return
	}

	book := repository.Book{
		Title:      req.Title,
		Author:     req.Author,
		IsBorrowed: false,
	}

	err = h.usecase.InsertBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when inserting book",
			RawErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
	})
}

func (h *Handler) GetBookByID(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:       http.StatusInternalServerError,
			ErrorMessage: "Need id params",
		})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when converting id to integer",
			RawErrorMessage: err.Error(),
		})
		return
	}

	book, err := h.usecase.FindBookByID(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when finding book",
			RawErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
		Data: presentation.BookResponse{
			Title:      book.Title,
			Author:     book.Author,
			IsBorrowed: book.IsBorrowed,
		},
	})
}

func (h *Handler) EditBook(ctx *gin.Context) {
	req := presentation.EditBookRequestBody{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil || req.ID <= 0 || req.Author == "" || req.Title == "" {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Given Request Data is not valid",
			RawErrorMessage: err.Error(),
		})
		return
	}

	book := repository.Book{
		ID:     req.ID,
		Author: req.Author,
		Title:  req.Title,
	}

	book, err = h.usecase.EditBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when editing book",
			RawErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
		Data: presentation.BookResponse{
			Title:      book.Title,
			Author:     book.Author,
			IsBorrowed: book.IsBorrowed,
		},
	})
}

func (h *Handler) DeleteBook(ctx *gin.Context) {
	req := presentation.DeleteBookRequestBody{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil || req.ID <= 0 {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Given Request Data is not valid",
			RawErrorMessage: err.Error(),
		})
		return
	}

	err = h.usecase.DeleteBookByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when deleting book",
			RawErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
	})
}

func (h *Handler) GetAllBook(ctx *gin.Context) {
	books, err := h.usecase.GetAllBook()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when retrieving books",
			RawErrorMessage: err.Error(),
		})
		return
	}
	var data []presentation.AllBookResponse
	for _, book := range books {
		data = append(data, presentation.AllBookResponse{
			Title:      book.Title,
			Author:     book.Author,
			IsBorrowed: book.IsBorrowed,
			ID:         book.ID,
		})
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
		Data:   data,
	})
}

func (h *Handler) BorrowBook(ctx *gin.Context) {
	username := "tasku"
	req := presentation.BorrowBookRequestBody{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil || req.BookID <= 0 {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Given Request Data is not valid",
			RawErrorMessage: err.Error(),
		})
		return
	}

	err = h.usecase.BorrowBook(username, req.BookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when borrowing book",
			RawErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
	})
}
