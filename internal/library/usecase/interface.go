package usecase

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type LibraryRepository interface {
	GetLibraryByID(libraryID int) (repo.Library, error)
	InsertLibrary(libraryID int) error
	UpdateLibrary(library repo.Library) (repo.Library, error)
	DeleteLibrary(libraryID int) error
}
