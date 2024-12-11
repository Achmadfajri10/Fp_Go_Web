package authmodel

import (
	"Fp_Go_Web/config"
	"Fp_Go_Web/entities"
)

func FindUserByUsername(username string) (entities.User, error) {
	var user entities.User
	result := config.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}

func FindUserByEmail(email string) (entities.User, error) {
	var user entities.User
	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}

func Create(user entities.User) bool {
	result := config.DB.Create(&user)
	return result.RowsAffected > 0
}
