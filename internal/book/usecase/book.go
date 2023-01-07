package usecase

import (
	"fmt"

	repo "github.com/arroganthooman/library-system/internal/repository"
)

func (u *Usecase) GetAllBook() ([]repo.Book, error) {
	books, err := u.repositories.GetAllBook()
	if err != nil {
		return []repo.Book{}, nil
	}

	return books, nil
}

func (u *Usecase) FindBookByID(bookID int) (repo.Book, error) {
	book, err := u.repositories.GetBookByID(bookID)
	if err != nil {
		return repo.Book{}, nil
	}

	return book, nil
}

func (u *Usecase) InsertBook(book repo.Book) error {
	err := u.repositories.InsertBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) EditBook(book repo.Book) (repo.Book, error) {
	book, err := u.repositories.UpdateBook(book)
	if err != nil {
		return repo.Book{}, nil
	}

	return book, nil
}

func (u *Usecase) DeleteBookByID(bookID int) error {
	err := u.repositories.DeleteBook(bookID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) BorrowBook(username string, bookID int) error {
	var book repo.Book
	book, err := u.repositories.GetBookByID(bookID)
	if err != nil {
		return fmt.Errorf("[Usecase][BorrowBook] error when calling GetBookByID, trace: %+v", err)
	}

	book.IsBorrowed = true
	book.UserUsername = username
	_, err = u.repositories.UpdateBook(book)
	if err != nil {
		return fmt.Errorf("[Usecase][UpdateBook] error when calling UpdateBook, trace: %+v", err)
	}

	return nil
}
