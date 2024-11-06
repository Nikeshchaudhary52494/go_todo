package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Nikeshchaudhary52494/goTest/models"
	"golang.org/x/crypto/bcrypt"
)

var users []models.User

func LoadUsers() error {
	data, err := os.ReadFile("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			users = []models.User{}
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &users)
}

func SaveUsers() error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("users.json", data, 0644)
}

func RegisterUser(username, password string) error {
	if GetUserByUsername(username) != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}
	users = append(users, user)
	return SaveUsers()
}

func GetUserByUsername(username string) *models.User {
	for _, user := range users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func ValidateUserPassword(username, password string) bool {
	user := GetUserByUsername(username)
	if user == nil {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
