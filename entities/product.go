package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category  	Category
	Stock       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}