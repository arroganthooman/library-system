package usecase

import (
	repo "github.com/arroganthooman/library-system/internal/repository"
)

type MockBookRepository struct {
	GetAllBookRes  []repo.Book
	GetAllBookErr  error
	GetBookByIDRes repo.Book
	GetBookByIDErr error
	InsertBookErr  error
	UpdateBookRes  repo.Book
	UpdateBookErr  error
	DeleteBookErr  error
}

func (m *MockBookRepository) GetAllBook() ([]repo.Book, error) {
	return m.GetAllBookRes, m.GetAllBookErr
}

func (m *MockBookRepository) GetBookByID(bookID int) (repo.Book, error) {
	return m.GetBookByIDRes, m.GetBookByIDErr
}

func (m *MockBookRepository) InsertBook(book repo.Book) error {
	return m.InsertBookErr
}

func (m *MockBookRepository) UpdateBook(book repo.Book) (repo.Book, error) {
	return m.UpdateBookRes, m.UpdateBookErr
}

func (m *MockBookRepository) DeleteBook(bookID int) error {
	return m.DeleteBookErr
}
