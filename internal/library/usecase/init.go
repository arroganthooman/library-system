package usecase

import (
	repository "github.com/arroganthooman/library-system/internal/repository"
)

type Usecase struct {
	repositories *repository.Repository
}

func NewLibraryRepo(repo *repository.Repository) *Usecase {
	return &Usecase{
		repositories: repo,
	}
}
