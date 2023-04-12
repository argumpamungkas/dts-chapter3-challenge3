package repository

import (
	"DTS/Chapter-3/chapter3-challenge3/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id uint) *entity.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		product := arguments.Get(0).(entity.Product)
		return &product
	}

}

func (repository *ProductRepositoryMock) FindAll() []*entity.Product {

	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).([]*entity.Product)
	return product

}
