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

func FindUserByID(ID uint) (entities.User, error) {
	var user entities.User
	result := config.DB.Where("ID = ?", ID).First(&user)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}

func Create(user entities.User) bool {
	result := config.DB.Create(&user)
	return result.RowsAffected > 0
}

func Update(id uint, user entities.User) error {
	var currUser entities.User
	err := config.DB.First(&currUser, id).Error
	if err != nil {
		return err
	}
	currUser.Username = user.Username
	currUser.Email = user.Email
	currUser.Password = user.Password
	result := config.DB.Save(&currUser)
	return result.Error
}

func Delete(id uint) error {
	var user entities.User
	result := config.DB.Delete(&user, id)
	return result.Error
}