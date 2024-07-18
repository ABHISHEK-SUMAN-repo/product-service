package repository

import (
	"product-service/config"
    "product-service/model"
)

func CreateProducts(products []model.Products) ([]model.Products, error) {

	result := config.DB.Save(&products)
	if result.Error != nil {
		return []model.Products{}, result.Error
	}
	return products, nil
}

func GetProductsByName(phoneNumber string) (model.Products, error) {
	var user model.Products
	result := config.DB.Where("phone = ?", phoneNumber).First(&user)
	if result.Error != nil {
		return model.Products{}, result.Error
	}
	return user, nil
}

func GetProductsByIds(ids []string) ([]model.Products, error) {
	var products []model.Products
	result := config.DB.Where("id IN ?", ids).Find(&products)
	if result.Error != nil {
		return []model.Products{}, result.Error
	}
	return products, nil
}

