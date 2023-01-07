package main

import (
	bookDeliv "github.com/arroganthooman/library-system/internal/book/delivery"
	bookUcase "github.com/arroganthooman/library-system/internal/book/usecase"
	repository "github.com/arroganthooman/library-system/internal/repository"
	userDeliv "github.com/arroganthooman/library-system/internal/user/delivery"
	userUcase "github.com/arroganthooman/library-system/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo, err := repository.InitRepo()
	if err != nil {
		panic(err)
	}

	// Init library
	userUsecase := userUcase.NewLibraryRepo(repo)
	userDelivery := userDeliv.NewLibraryHandler(router, userUsecase)
	userDelivery.SetEndpoint()

	// Init book
	bookUsecase := bookUcase.NewBookRepo(repo)
	bookDelivery := bookDeliv.NewBookHandler(router, bookUsecase)
	bookDelivery.SetEndpoint()

	router.Run(":8080")
}
