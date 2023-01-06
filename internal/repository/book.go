package repository

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Title     string
	Author    string
	Library   Library `gorm:"foreignKey:LibraryID"`
	LibraryID int
}

// SQL REPO
func (repo *Repository) GetBookByID(bookID int) (Book, error) {
	return Book{}, nil
}

func (repo *Repository) InsertBook(bookID int) error {
	return nil
}

func (repo *Repository) UpdateBook(book Book) (Book, error) {
	return Book{}, nil
}

func (repo *Repository) DeleteBook(bookID int) error {
	return nil
}

// Redis REPO
