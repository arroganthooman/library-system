package delivery

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type Usecase interface {
	FindUserByUsername(username string) (repo.User, error)
	InsertUser(repo.User) error
	EditUser(user repo.User) (repo.User, error)
	DeleteUserByUsername(username string) error
}
