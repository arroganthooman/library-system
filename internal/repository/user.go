package repository

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
	Books    []Book `json:"borrowed_books"` //Borrowed book
}

type AuthToken struct {
	gorm.Model
	Token       string    `json:"token"`
	ExpiredDate time.Time `json:expired_date`
	Username    string
}

// SQL REPO
func (repo *Repository) GetUser(username string) (User, error) {
	var user User
	res := repo.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		return User{}, res.Error
	} else if res.RowsAffected == int64(0) {
		return User{}, fmt.Errorf("[Repo][GetUser] No record found")
	}

	return user, nil
}

func (repo *Repository) InsertUser(username string, password string) error {
	var user User
	res := repo.DB.First(&user, "username = ?", username)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return res.Error
	} else if user.Username == username {
		return fmt.Errorf("[Repo][InsertUser] Username already exists")
	}

	user = User{
		Username: username,
		Password: password,
	}

	err := repo.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateUser(user User) (User, error) {
	err := repo.DB.Save(user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (repo *Repository) DeleteUser(username string) error {
	err := repo.DB.Exec("DELETE FROM users where username = ?", username).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) CreateToken(username string) (string, error) {
	expired := time.Now().Add(time.Hour * 48)
	textToEncode := fmt.Sprintf("%s:%d", username, expired)
	sha := sha1.New()
	sha.Write([]byte(textToEncode))

	encrypted := sha.Sum(nil)
	token := fmt.Sprintf("%x", encrypted)

	tokenObject := AuthToken{
		Token:       token,
		ExpiredDate: expired,
		Username:    username,
	}

	err := repo.DB.Create(&tokenObject).Error
	if err != nil {
		return "", err
	}

	redisRes := repo.Redis.Set(context.Background(), fmt.Sprintf("token:%s", token), token, time.Hour*48)
	if redisRes.Err() != nil {
		fmt.Println("[Repo][CreateToken] Failed to insert to redis")
	}

	return token, nil
}

func (r *Repository) UserBorrowBook(username string, bookID int) error {
	var user User
	err := r.DB.First(&user, "username = ?", "username").Error
	if err != nil {
		return err
	}

	var book Book
	err = r.DB.First(&book, "id = ?", bookID).Error
	if err != nil {
		return err
	}

	book.IsBorrowed = true
	book.UserUsername = user.Username
	err = r.DB.Save(&book).Error
	if err != nil {
		return err
	}

	return nil
}
