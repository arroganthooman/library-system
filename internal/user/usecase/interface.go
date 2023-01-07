package usecase

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type UserRepository interface {
	GetUser(username string) (repo.User, error)
	InsertUser(username string, password string) error
	UpdateUser(user repo.User, oldUsername string) (repo.User, error)
	DeleteUser(username string) error
	CreateToken(username string) (string, error)
	GetBorrowedBooksByUsername(username string) ([]repo.Book, error)
}
