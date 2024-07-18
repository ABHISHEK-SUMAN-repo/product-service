package adapter

import (
	"product-service/dto"
	"product-service/model"
	"time"
)

func ConvertProduct(productDTO dto.ProductDTO) model.Products {
	return model.Products{
		ProductName: productDTO.ProductName,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Quantity:    productDTO.Quantity,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
		Created_by:  productDTO.Created_by,
		Updated_by:  productDTO.Updated_by,
	}
}
