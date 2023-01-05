package usecase

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type BookRepository interface {
	GetBookByID(bookID int) (repo.Book, error)
	InsertBook(bookID int) error
	UpdateBook(book repo.Book) (repo.Book, error)
	DeleteBook(bookID int) error
}