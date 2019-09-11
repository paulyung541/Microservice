package model

// User 用户
type User struct {
	ID       int
	Name     string
	Account  string
	Password string
}

func (User) TableName() string {
	return "user"
}
