package repository

import (
	"DTS/Chapter-3/chapter3-challenge3/entity"
)

type ProductRepository interface {
	FindById(id uint) *entity.Product
	FindAll() []*entity.Product
}
