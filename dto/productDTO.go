package dto

import (
	"time"
)

type ProductDTO struct {
	ProductName *string   `json:"productName"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price"`
	Quantity    *float64  `json:"quantity"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Created_by  *string   `json:"created_by"`
	Updated_by  *string   `json:"updated_by"`
}
