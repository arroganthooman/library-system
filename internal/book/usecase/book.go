package usecase

import (
	"fmt"
	"log"

	repo "github.com/arroganthooman/library-system/internal/repository"
)

func (u *Usecase) GetAllBook() ([]repo.Book, error) {
	books, err := u.repositories.GetAllBook()
	if err != nil {
		log.Printf("[Usecase][GetAllBook] err when calling GetAllBook: %+v", err)
		return []repo.Book{}, nil
	}

	return books, nil
}

func (u *Usecase) FindBookByID(bookID int) (repo.Book, error) {
	book, err := u.repositories.GetBookByID(bookID)
	if err != nil {
		log.Printf("[Usecase][FindBookByID] err when calling GetBookByID: %+v", err)
		return repo.Book{}, nil
	}

	return book, nil
}

func (u *Usecase) InsertBook(book repo.Book) error {
	err := u.repositories.InsertBook(book)
	if err != nil {
		log.Printf("[Usecase][InsertBook] err when calling InsertBook: %+v", err)
		return err
	}

	return nil
}

func (u *Usecase) EditBook(book repo.Book) (repo.Book, error) {
	book, err := u.repositories.UpdateBook(book)
	if err != nil {
		log.Printf("[Usecase][EditBook] err when calling UpdateBook: %+v", err)
		return repo.Book{}, nil
	}

	return book, nil
}

func (u *Usecase) DeleteBookByID(bookID int) error {
	err := u.repositories.DeleteBook(bookID)
	if err != nil {
		log.Printf("[Usecase][DeleteBookByID] err when calling DeleteBook: %+v", err)
		return err
	}

	return nil
}

func (u *Usecase) BorrowBook(username string, bookID int) error {
	var book repo.Book
	book, err := u.repositories.GetBookByID(bookID)
	if err != nil {
		log.Printf("[Usecase][BorrowBook] err when calling GetBookByID: %+v", err)
		return fmt.Errorf("[Usecase][BorrowBook] error when calling GetBookByID, trace: %+v", err)
	}

	if book.IsBorrowed {
		return fmt.Errorf("[Usecase][BorrowBook] You cannot borrow this book, it's borrowed by other")
	}

	book.IsBorrowed = true
	book.UserUsername = username
	_, err = u.repositories.UpdateBook(book)
	if err != nil {
		log.Printf("[Usecase][BorrowBook] err when calling UpdateBook: %+v", err)
		return fmt.Errorf("[Usecase][UpdateBook] error when calling UpdateBook, trace: %+v", err)
	}

	return nil
}

func (u *Usecase) ReturnBook(username string, bookID int) error {
	var book repo.Book
	book, err := u.repositories.GetBookByID(bookID)
	if err != nil {
		log.Printf("[Usecase][ReturnBook] err when calling GetBookByID: %+v", err)
		return fmt.Errorf("[Usecase][ReturnBook] error when calling GetBookByID")
	}

	if book.UserUsername != username {
		return fmt.Errorf("[Usecase][ReturnBook] You are not borrowing this book")
	}

	book.IsBorrowed = false
	book.UserUsername = ""

	_, err = u.repositories.UpdateBook(book)
	if err != nil {
		log.Printf("[Usecase][ReturnBook] err when calling UpdateBook: %+v", err)
		return fmt.Errorf("[Usecase][ReturnBook] error when calling UpdateBook")
	}

	return nil
}
