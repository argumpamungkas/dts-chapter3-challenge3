package service

import (
	"DTS/Chapter-3/chapter3-challenge3/entity"
	"DTS/Chapter-3/chapter3-challenge3/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id uint) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found")
	} else {
		return product, nil
	}
}

func (service ProductService) GetAll() ([]*entity.Product, error) {
	product := service.Repository.FindAll()
	if product == nil {
		return nil, errors.New("All product not found")
	} else {
		return product, nil
	}
}
