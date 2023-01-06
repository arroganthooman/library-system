package delivery

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type Usecase interface {
	FindBookByID(bookID int) (repo.Book, error)
	InsertBook(repo.Book) error
	EditBookByID(book repo.Book) (repo.Book, error)
	DeleteBookByID(bookID int) error
}
