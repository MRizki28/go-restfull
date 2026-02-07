package repository

import (
	"go-restfull/app/config"
	"go-restfull/app/model"
)

type ProductRepository interface {
	GetAllProducts(search string) ([]model.ProductModel, error)
	CreateDataProduct(product *model.ProductModel) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetAllProducts(search string) ([]model.ProductModel, error) {
	var products []model.ProductModel
	query := config.DB

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	result := query.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *productRepository) CreateDataProduct(product *model.ProductModel) error {
	result := config.DB.Create(product)
	return result.Error
}
