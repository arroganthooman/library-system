package repository

import (
	"fmt"
	"log"

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
		log.Printf("[Repo][GetAllBook] err: %+v", err)
		return nil, err
	}

	return books, nil
}

func (repo *Repository) GetBookByID(bookID int) (Book, error) {
	var book Book
	res := repo.DB.First(&book, "ID = ?", bookID)

	if res.Error != nil {
		log.Printf("[Repo][GetBookByID] err: %+v", res.Error)
		return Book{}, res.Error
	} else if res.RowsAffected == int64(0) {
		log.Print("[Repo][GetBookByID] err: No record found")
		return Book{}, fmt.Errorf("[Repo][GetBookByID] No record found")
	}

	return book, nil
}

func (repo *Repository) GetBorrowedBooksByUsername(username string) ([]Book, error) {
	var books []Book
	res := repo.DB.Where("user_username = ?", username).Find(&books)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		log.Printf("[Repo][GetBorrowedBooksByUsername] err: %+v", res.Error)
		return []Book{}, res.Error
	}

	return books, nil
}

func (repo *Repository) InsertBook(book Book) error {
	err := repo.DB.Create(&book).Error
	if err != nil {
		log.Printf("[Repo][InsertBook] err: %+v", err)
		return err
	}

	return nil
}

func (repo *Repository) UpdateBook(book Book) (Book, error) {
	err := repo.DB.Save(book).Error
	if err != nil {
		log.Printf("[Repo][UpdateBook] err: %+v", err)
		return Book{}, err
	}

	return book, nil
}

func (repo *Repository) DeleteBook(bookID int) error {
	err := repo.DB.Exec("DELETE FROM books where id = ?", bookID).Error
	if err != nil {
		log.Printf("[Repo][DeleteBook] err: %+v", err)
		return err
	}

	return nil
}
