package main

import (
	bookDeliv "github.com/arroganthooman/library-system/internal/book/delivery"
	bookUcase "github.com/arroganthooman/library-system/internal/book/usecase"
	libDelivery "github.com/arroganthooman/library-system/internal/library/delivery"
	libUsecase "github.com/arroganthooman/library-system/internal/library/usecase"
	repository "github.com/arroganthooman/library-system/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo, err := repository.InitRepo()
	if err != nil {
		panic(err)
	}

	// Init library
	libraryUsecase := libUsecase.NewLibraryRepo(repo)
	libraryDelivery := libDelivery.NewLibraryHandler(router, libraryUsecase)
	libraryDelivery.SetEndpoint()

	// Init book
	bookUsecase := bookUcase.NewBookRepo(repo)
	bookDelivery := bookDeliv.NewBookHandler(router, bookUsecase)
	bookDelivery.SetEndpoint()

	router.Run(":8080")
}
