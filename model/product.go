package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type Products struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"_id"`
	ProductName *string   `json:"productName"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price"`
	Quantity    *float64  `json:"quantity"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Created_by  *string   `json:"created_by"`
	Updated_by  *string   `json:"updated_by"`
}

func (Products) TableName() string {
	return viper.GetString("products_table_name")
}
