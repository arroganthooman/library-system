package delivery

import (
	"net/http"

	"github.com/arroganthooman/library-system/internal/repository"
	"github.com/arroganthooman/library-system/presentation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	requestBody := presentation.AuthRequestBody{}
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Given Request Data is not valid",
			RawErrorMessage: err.Error(),
		})
		return
	}

	token, err := h.usecase.Login(requestBody.Username, requestBody.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			RawErrorMessage: err.Error(),
			ErrorMessage:    "Error when creating token",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, presentation.ResponseStandard{
			Status: http.StatusOK,
			Data: presentation.LoginResponse{
				Token: token,
			},
		})
		return
	}
}

func (h *Handler) RegisterUser(ctx *gin.Context) {
	requestBody := presentation.AuthRequestBody{}
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil || requestBody.Username == "" || requestBody.Password == "" {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Given Request Data is not valid",
			RawErrorMessage: err.Error(),
		})
		return
	}

	user := repository.User{
		Username: requestBody.Username,
		Password: requestBody.Password,
	}
	err = h.usecase.InsertUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when registering user",
			RawErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
	})

	return
}

func (h *Handler) EditUser(ctx *gin.Context) {
	requestBody := presentation.AuthRequestBody{}
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil || requestBody.Username == "" || requestBody.Password == "" {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Given Request Data is not valid. Need username and password",
			RawErrorMessage: err.Error(),
		})
		return
	}

	user := repository.User{
		Username: requestBody.Username,
		Password: requestBody.Password,
	}
	uname, isExists := ctx.Get("username")
	if !isExists {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:       http.StatusInternalServerError,
			ErrorMessage: "Error when getting username from context",
		})
		return
	}
	usernameFromAuth, ok := uname.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:       http.StatusInternalServerError,
			ErrorMessage: "Error when asserting type",
		})
		return
	}
	_, err = h.usecase.EditUser(user, usernameFromAuth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when editing user",
			RawErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
	})

	return
}

func (h *Handler) GetUserInfoFromAuthToken(ctx *gin.Context) {
	uname, isExists := ctx.Get("username")
	if !isExists {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:       http.StatusInternalServerError,
			ErrorMessage: "Error when getting username from context",
		})
		return
	}
	usernameFromAuth, ok := uname.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:       http.StatusInternalServerError,
			ErrorMessage: "Error when asserting type",
		})
		return
	}

	user, err := h.usecase.GetUserInfo(usernameFromAuth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presentation.ResponseStandard{
			Status:          http.StatusInternalServerError,
			ErrorMessage:    "Error when getting user info",
			RawErrorMessage: err.Error(),
		})
		return
	}

	userResponse := presentation.UserInfoResponse{
		Username:      user.Username,
		BorrowedBooks: make([]presentation.BookResponse, 0),
	}
	for _, book := range user.Books {
		userResponse.BorrowedBooks = append(userResponse.BorrowedBooks, presentation.BookResponse{
			Author:     book.Author,
			Title:      book.Title,
			ID:         book.ID,
			IsBorrowed: book.IsBorrowed,
		})
	}

	ctx.JSON(http.StatusOK, presentation.ResponseStandard{
		Status: http.StatusOK,
		Data:   userResponse,
	})

	return
}
