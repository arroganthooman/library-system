package delivery

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type Usecase interface {
	FindLibraryByID(libraryID int) (repo.Library, error)
	InsertLibrary(repo.Library) error
	EditLibraryByID(library repo.Library) (repo.Library, error)
	DeleteLibraryByID(libraryID int) error
}
