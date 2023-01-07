package usecase

import (
	"fmt"

	repo "github.com/arroganthooman/library-system/internal/repository"
)

func (u *Usecase) FindUserByUsername(username string) (repo.User, error) {
	user, err := u.repositories.GetUser(username)
	if err != nil {
		return repo.User{}, err
	}

	return user, nil
}

func (u *Usecase) InsertUser(user repo.User) error {
	err := u.repositories.InsertUser(user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) EditUser(user repo.User) (repo.User, error) {
	return repo.User{}, nil
}

func (u *Usecase) DeleteUserByUsername(username string) error {
	return nil
}

func (u *Usecase) Login(username string, password string) (token string, err error) {
	res, err := u.repositories.GetUser(username)
	if err != nil {
		return "", err
	}

	if res.Password != password {
		return "", fmt.Errorf("[Usecase][Login] Invalid Password")
	}

	token, err = u.repositories.CreateToken(username)
	if err != nil {
		return "", err
	}

	return token, err
}
