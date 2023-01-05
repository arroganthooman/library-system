package repository

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func InitRepo() (*Repository, error) {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		return &Repository{}, nil
	}

	db.AutoMigrate(&Book{}, &Library{})

	return &Repository{
		DB: db,
	}, nil
}
