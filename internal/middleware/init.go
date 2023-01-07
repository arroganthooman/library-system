package middleware

import (
	repository "github.com/arroganthooman/library-system/internal/repository"
)

type Middleware struct {
	repositories *repository.Repository
}

func InitMiddleware(repo *repository.Repository) *Middleware {
	return &Middleware{
		repositories: repo,
	}
}
