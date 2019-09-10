package services

import (
	"Microservice/auth/config"
	"Microservice/auth/model"
)

// User 业务层
type User struct{}

// AddUser xxx
func (*User) AddUser(user *model.User) error {
	return config.Db.Create(user).Error
}

// GetUser xxx
func (*User) GetUser(account, password string) (*model.User, error) {
	var user model.User
	return &user, config.Db.Where("account = ? AND password = ?", account, password).First(&user).Error
}
