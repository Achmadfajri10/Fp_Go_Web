package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	Stock       int
	Description string
}
