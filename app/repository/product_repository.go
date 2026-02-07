package repository

import (
	"errors"
	"go-restfull/app/config"
	"go-restfull/app/model"
)

type ProductRepository interface {
	GetAllProducts(search string) ([]model.ProductModel, error)
	CreateDataProduct(product *model.ProductModel) error
	GetDataById(id int) (*model.ProductModel, error)
	UpdateData(id int, data *model.ProductModel) (*model.ProductModel, error)
	DeleteData(id int) error
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

func (r *productRepository) GetDataById(id int) (*model.ProductModel, error) {
	var product model.ProductModel
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *productRepository) UpdateData(id int,data *model.ProductModel,) (*model.ProductModel, error) {

	var product model.ProductModel

	if err := config.DB.First(&product, id).Error; err != nil {
		return nil, errors.New("product not found")
	}

	tx := config.DB.Model(&product).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("no data changed")
	}

	return &product, nil
}

func (r *productRepository) DeleteData(id int) error {
	var product model.ProductModel

	if err := config.DB.First(&product, id).Error; err != nil {
		return errors.New("product not found")
	}

	tx := config.DB.Delete(&product)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("no data deleted")
	}

	return nil
}
