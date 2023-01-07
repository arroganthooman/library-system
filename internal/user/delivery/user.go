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
