package usecase

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type BookRepository interface {
	GetAllBook() ([]repo.Book, error)
	GetBookByID(bookID int) (repo.Book, error)
	InsertBook(book repo.Book) error
	UpdateBook(book repo.Book) (repo.Book, error)
	DeleteBook(bookID int) error
}
