package service

import (
	"go-restfull/app/model"
	"go-restfull/app/repository"
	"go-restfull/app/request"
)

type ProductService interface {
	GetAllProducts(search string) ([]model.ProductModel, error)
	CreateDataProduct(req request.CreateProductRequest) (product *model.ProductModel, err error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts(search string) ([]model.ProductModel, error) {
	return s.repo.GetAllProducts(search)
}

func (s *productService) CreateDataProduct(req request.CreateProductRequest) (product *model.ProductModel, err error) {
	product = &model.ProductModel{
		Name:        req.Name,
		Price:       req.Price,
	}

	if err := s.repo.CreateDataProduct(product); err != nil {
		return nil, err
	}
	
	return product, nil
}


