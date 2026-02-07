package model

import "time"

type ProductModel struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ProductModel) TableName() string {
	return "tb_product"
}