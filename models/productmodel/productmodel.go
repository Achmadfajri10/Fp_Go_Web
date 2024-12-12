package productmodel

import (
	"Fp_Go_Web/config"
	"Fp_Go_Web/entities"
)

func GetAll() ([]entities.Product, error) {
	var products []entities.Product
	err := config.DB.Preload("Category").Find(&products).Error
	return products, err
}

func Create(product entities.Product) bool {
	result := config.DB.Create(&product)
	return result.Error == nil
}

func Detail(id uint) (entities.Product, error) {
	var product entities.Product
	err := config.DB.Preload("Category").First(&product, id).Error
	return product, err
}

func Update(id uint, product entities.Product) error {
	var currProduct entities.Product
	err := config.DB.First(&currProduct, id).Error
	if err != nil {
		return err
	}
	currProduct.Name = product.Name
	currProduct.CategoryID = product.CategoryID
	currProduct.Stock = product.Stock
	currProduct.Description = product.Description
	result := config.DB.Save(&currProduct)
	return result.Error

}

func Delete(id uint) error {
	result := config.DB.Delete(&entities.Product{}, id)
	return result.Error
}
