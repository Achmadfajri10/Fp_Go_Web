package categorymodel

import (
	"Fp_Go_Web/config"
	"Fp_Go_Web/entities"
)

func GetAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func Create(category entities.Category) bool {
	result := config.DB.Create(&category)
	return result.RowsAffected > 0
}

func Detail(id uint) (entities.Category, error) {
	var category entities.Category
	err := config.DB.First(&category, id).Error
	return category, err
}

func Update(id uint, category entities.Category) bool {
	result := config.DB.Model(&category).Where("id = ?", id).Updates(category)
	return result.RowsAffected > 0
}

func Delete(id uint) error { // Changed id to uint
	result := config.DB.Delete(&entities.Category{}, id) // GORM's Delete method
	return result.Error
}
