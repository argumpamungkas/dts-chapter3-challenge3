package service

import (
	"DTS/Chapter-3/chapter3-challenge3/entity"
	"DTS/Chapter-3/chapter3-challenge3/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepo}

func TestGetOneProduct_NotFound(t *testing.T) {

	productRepo.Mock.On("FindById", uint(1)).Return(nil)

	product, err := productService.GetOneProduct(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found", err.Error())

}

func TestGetOneProduct_Found(t *testing.T) {
	product := entity.Product{
		ID:          uint(2),
		Title:       "Handphone",
		Description: "Samsung J5",
	}

	productRepo.Mock.On("FindById", uint(2)).Return(product)

	result, err := productService.GetOneProduct(2)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.ID, result.ID)
	assert.Equal(t, product.Title, result.Title)
}

func TestGetAllProduct_NotFound(t *testing.T) {

	productRepo.Mock.On("FindAll").Return(nil)

	result, err := productService.GetAll()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "All product not found", err.Error())
}

func TestGetAllProduct_Found(t *testing.T) {

	product := []*entity.Product{
		{ID: uint(4), Title: "Handphone", Description: "Samsung J5"},
		{ID: uint(5), Title: "Laptop", Description: "Lenovo Slim"},
		{ID: uint(6), Title: "TWS", Description: "Airbuds Robot"},
	}

	productRepo.Mock.On("FindAll").Return(product)

	res, err := productService.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, product, res)
}
