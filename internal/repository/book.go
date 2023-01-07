package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID           int `gorm:"primaryKey"`
	Title        string
	Author       string
	IsBorrowed   bool
	UserUsername string
}

// SQL REPO

func (repo *Repository) GetAllBook() ([]Book, error) {
	var books []Book
	err := repo.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (repo *Repository) GetBookByID(bookID int) (Book, error) {
	var book Book
	res := repo.DB.First(&book, "ID = ?", bookID)

	if res.Error != nil {
		return Book{}, res.Error
	} else if res.RowsAffected == int64(0) {
		return Book{}, fmt.Errorf("[Repo][GetBookByID] No record found")
	}

	return book, nil
}

func (repo *Repository) InsertBook(book Book) error {
	err := repo.DB.Create(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateBook(book Book) (Book, error) {
	err := repo.DB.Save(book).Error
	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (repo *Repository) DeleteBook(bookID int) error {
	err := repo.DB.Exec("DELETE FROM books where id = ?", bookID).Error
	if err != nil {
		return err
	}

	return nil
}
