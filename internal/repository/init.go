package repository

import (
	"github.com/glebarez/sqlite"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Repository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func InitRepo() (*Repository, error) {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		return &Repository{}, nil
	}
	db.AutoMigrate(&Book{}, &Library{})

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &Repository{
		DB:    db,
		Redis: rdb,
	}, nil
}
