package services

import (
	"github.com/paulyung541/jotnar"
	"Microservice/auth/model"
)

// User 业务层
type User struct{}

// AddUser xxx
func (*User) AddUser(user *model.User) error {
	return jotnar.WriteGorm().Create(user).Error
}

// GetUser xxx
func (*User) GetUser(account, password string) (*model.User, error) {
	var user model.User
	return &user, jotnar.ReadGorm().Where("account = ? AND password = ?", account, password).First(&user).Error
}
