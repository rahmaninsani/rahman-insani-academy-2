package service

import (
	"01-soal-1/entity"
	"errors"
)

type ProductService struct {
	data []entity.Product
}

func (service *ProductService) FindAll() ([]entity.Product, error) {
	if service.data == nil {
		return nil, errors.New("products not found")
	}

	return service.data, nil

}

func (service *ProductService) FindById(productId uint64) (entity.Product, error) {
	var product entity.Product

	if service.data == nil {
		return product, errors.New("product not found")
	}

	for _, value := range service.data {
		if value.Id == productId {
			product = value
			break
		}
	}
	return product, nil
}
