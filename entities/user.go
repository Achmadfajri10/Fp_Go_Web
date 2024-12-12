package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

type LoginInput struct {
	Identifier string
	Password   string
}

type EditInput struct {
	ID          uint
	Username    string
	Email       string
	OldPassword string
	Password    string
}
