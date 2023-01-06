package usecase

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

func (u *Usecase) FindBookByID(bookID int) (repo.Book, error) {
	return repo.Book{}, nil
}

func (u *Usecase) InsertBook(repo.Book) error {
	return nil
}

func (u *Usecase) EditBookByID(book repo.Book) (repo.Book, error) {
	return repo.Book{}, nil
}

func (u *Usecase) DeleteBookByID(bookID int) error {
	return nil
}
