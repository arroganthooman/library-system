package repository

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/arroganthooman/library-system/presentation"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primaryKey;unique"`
	Password string `json:"password"`
	Books    []Book `json:"borrowed_books"` //Borrowed book
}

type AuthToken struct {
	gorm.Model
	Token       string    `gorm:"primaryKey" json:"token"`
	ExpiredDate time.Time `json:"expired_date"`
	Username    string
}

// SQL REPO
func (repo *Repository) GetUser(username string) (User, error) {
	var user User
	res := repo.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		log.Printf("[Repo][GetUser] err: %+v", res.Error)
		return User{}, res.Error
	} else if res.RowsAffected == int64(0) {
		log.Printf("[Repo][GetUser] err: No record found")
		return User{}, fmt.Errorf("[Repo][GetUser] No record found")
	}

	return user, nil
}

func (repo *Repository) InsertUser(username string, password string) error {
	var user User
	res := repo.DB.First(&user, "username = ?", username)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		log.Printf("[Repo][InsertUser] err: %+v", res.Error)
		return res.Error
	} else if user.Username == username {
		log.Printf("[Repo][InsertUser] err: %+v", "Username already exists")
		return fmt.Errorf("[Repo][InsertUser] Username already exists")
	}

	user = User{
		Username: username,
		Password: password,
	}

	err := repo.DB.Create(&user).Error
	if err != nil {
		log.Printf("[Repo][InsertUser] err: %+v", err)
		return err
	}

	return nil
}

func (repo *Repository) UpdateUser(user User, oldUsername string) (User, error) {
	var oldUser User
	err := repo.DB.First(&oldUser, "username = ?", oldUsername).Error
	if err != nil {
		log.Printf("[Repo][UpdateUser] err: %+v", err)
		return User{}, err
	}

	updatedUser := User{}
	if user.Username != "" {
		updatedUser.Username = user.Username
	}

	if user.Password != "" {
		updatedUser.Password = user.Password
	}

	err = repo.DB.Model(&oldUser).Updates(updatedUser).Error
	if err != nil {
		log.Printf("[Repo][UpdateUser] err: %+v", err)
		return User{}, err
	}

	return user, nil
}

func (repo *Repository) DeleteUser(username string) error {
	err := repo.DB.Exec("DELETE FROM users where username = ?", username).Error
	if err != nil {
		log.Printf("[Repo][UpdateUser] err: %+v", err)
		return err
	}

	return nil
}

func (repo *Repository) CreateToken(username string) (string, error) {
	expired := time.Now().Add(time.Hour * 48)
	textToEncode := fmt.Sprintf("%s:%+v", username, expired)
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
		log.Printf("[Repo][CreateToken] err: %+v", err)
		return "", err
	}

	redisRes := repo.Redis.Set(context.Background(), fmt.Sprintf("token:%s", token), token, time.Hour*48)
	if redisRes.Err() != nil {
		log.Printf("[Repo][CreateToken] err: %+v", "Failed to insert to redis")
		fmt.Println("[Repo][CreateToken] Failed to insert to redis")
	}

	redisRes = repo.Redis.Set(context.Background(), fmt.Sprintf("token_username:%s", username), token, time.Hour*48)
	if redisRes.Err() != nil {
		log.Printf("[Repo][CreateToken] err: %+v", "Failed to insert to redis")
		fmt.Println("[Repo][CreateToken] Failed to insert to redis")
	}

	return token, nil
}

func (repo *Repository) CheckToken(token string) (string, error) {
	redisRes := repo.Redis.Get(context.Background(), fmt.Sprintf("token:%s", token))
	if redisRes.Err() != nil {
		log.Printf("[Repo][CheckToken] err: %+v", redisRes.Err())
	} else {
		redisUnameRes := repo.Redis.Get(context.Background(), fmt.Sprintf("token_username:%s", redisRes))
		if redisUnameRes.Err() == nil {
			return redisUnameRes.String(), nil
		}
	}

	var authToken AuthToken
	authToken.Token = token
	err := repo.DB.First(&authToken).Error
	if err != nil {
		log.Printf("[Repo][CheckToken] err: %+v", err)
		return "", fmt.Errorf("[Repo][CheckToken] Error when checking token, trace: %+v", err)
	}

	if authToken.ExpiredDate.Before(time.Now()) {
		return "", presentation.ErrorUnauthorized
	}

	return authToken.Username, nil
}
