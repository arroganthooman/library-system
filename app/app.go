package main

import (
	bookDeliv "github.com/arroganthooman/library-system/internal/book/delivery"
	bookUcase "github.com/arroganthooman/library-system/internal/book/usecase"
	mware "github.com/arroganthooman/library-system/internal/middleware"
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

	// Init middleware
	middleware := mware.InitMiddleware(repo)

	// Init library
	userUsecase := userUcase.NewLibraryUsecase(repo)
	userDelivery := userDeliv.NewLibraryHandler(router, userUsecase, middleware)
	userDelivery.SetEndpoint()

	// Init book
	bookUsecase := bookUcase.NewBookUsecase(repo)
	bookDelivery := bookDeliv.NewBookHandler(router, bookUsecase, middleware)
	bookDelivery.SetEndpoint()

	router.Run(":8080")
}
