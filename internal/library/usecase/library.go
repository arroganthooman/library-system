package usecase

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

func (u *Usecase) FindLibraryByID(libraryID int) (repo.Library, error) {
	return repo.Library{}, nil
}

func (u *Usecase) InsertLibrary(repo.Library) error {
	return nil
}

func (u *Usecase) EditLibraryByID(library repo.Library) (repo.Library, error) {
	return repo.Library{}, nil
}

func (u *Usecase) DeleteLibraryByID(libraryID int) error {
	return nil
}
