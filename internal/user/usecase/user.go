package usecase

import (
	"fmt"
	"log"

	repo "github.com/arroganthooman/library-system/internal/repository"
)

func (u *Usecase) FindUserByUsername(username string) (repo.User, error) {
	user, err := u.repositories.GetUser(username)
	if err != nil {
		log.Printf("[Usecase][FindUserByUsername] error when calling GetUser: %+v", err)
		return repo.User{}, err
	}

	return user, nil
}

func (u *Usecase) InsertUser(user repo.User) error {
	err := u.repositories.InsertUser(user.Username, user.Password)
	if err != nil {
		log.Printf("[Usecase][InsertUser] error when calling InsertUser: %+v", err)
		return err
	}

	return nil
}

func (u *Usecase) EditUser(user repo.User, usernameFromAuth string) (repo.User, error) {
	userFromDB, err := u.repositories.GetUser(user.Username)
	if err != nil {
		log.Printf("[Usecase][EditUser] error when calling GetUser: %+v", err)
		return repo.User{}, fmt.Errorf("[Usecase][EditUser] error when calling to GetUser, trace: %+v", err)
	}

	if userFromDB.Username != usernameFromAuth {
		return repo.User{}, fmt.Errorf("[Usecase][EditUser] you're not editing your account based on auth token")
	}

	res, err := u.repositories.UpdateUser(user, usernameFromAuth)
	if err != nil {
		log.Printf("[Usecase][EditUser] error when calling UpdateUser: %+v", err)
		return repo.User{}, fmt.Errorf("[Usecase][EditUser] error when calling to UpdateUser, trace: %+v", err)
	}

	return res, nil

}

func (u *Usecase) Login(username string, password string) (token string, err error) {
	res, err := u.repositories.GetUser(username)
	if err != nil {
		log.Printf("[Usecase][Login] err when calling GetUser: %+v", err)
		return "", fmt.Errorf("[Usecase][Login] error when calling to GetUser, trace: %+v", err)
	}

	if res.Password != password {
		return "", fmt.Errorf("[Usecase][Login] Invalid Password")
	}

	token, err = u.repositories.CreateToken(username)
	if err != nil {
		log.Printf("[Usecase][Login] error when calling CreateToken: %+v", err)
		return "", err
	}

	return token, err
}

func (u *Usecase) GetUserInfo(username string) (user repo.User, err error) {
	user, err = u.FindUserByUsername(username)
	if err != nil {
		log.Printf("[Usecase][GetUserInfo] error when calling FindUserByUsername: %+v", err)
		return repo.User{}, fmt.Errorf("[Usecase][GetUserInfo] error when calling to FindUserByUsername, trace: %+v", err)
	}

	borrowedBooks, err := u.repositories.GetBorrowedBooksByUsername(user.Username)
	if err != nil {
		log.Printf("[Usecase][GetUserInfo] error when calling GetBorrowedBooksByUsername: %+v", err)
		return repo.User{}, fmt.Errorf("[Usecase][GetUserInfo] error when calling to GetBorrowedBooksByUsername, trace: %+v", err)
	}

	user.Books = borrowedBooks
	return user, nil
}
