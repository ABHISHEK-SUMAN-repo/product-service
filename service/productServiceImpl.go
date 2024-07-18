package service

import (
	"product-service/adapter"
	"product-service/dto"
	"product-service/model"
	"product-service/repository"
)

func CreateProducts(productDTO []dto.ProductDTO) ([]model.Products, error) {

	var products []model.Products

	for _, p := range productDTO {
		product := adapter.ConvertProduct(p)
		products = append(products, product)
	}

	result,err := repository.CreateProducts(products)

	    if err!= nil {
            return nil, err
        }

	return result, nil
}


func GetProductsByIds(ids []string) ([]model.Products , error) {

	result, err := repository.GetProductsByIds(ids)
	    if err!= nil {
            return nil, err
        }
		
    return result, nil
}