package model

import (
	"06-golang-advance-2/helper"
	"errors"
)

var (
	listData []*Product
)

func init() {
	listData = []*Product{}
}

type Product struct {
	ID   int
	Name string
}

type ProductModel struct{}

func (model *ProductModel) FindAll() ([]*Product, error) {
	if len(listData) == 0 {
		return nil, errors.New("no record on database")
	}
	return listData, nil
}

func (model *ProductModel) FindOne(product *Product) (Product, error) {
	productId := -1

	for index, val := range listData {
		if val.ID == product.ID {
			productId = index
		}
	}

	if productId == -1 {
		return Product{}, errors.New("no data found")
	}

	return *listData[productId], nil
}

func (model *ProductModel) Insert(newProduct *Product) (Product, error) {
	if len(listData) == 0 {
		newProduct.ID = helper.GenerateID(0)
	} else {
		newProduct.ID = helper.GenerateID(listData[len(listData)-1].ID)
	}

	listData = append(listData, newProduct)

	return *newProduct, nil
}

func (model *ProductModel) Update(updatedUser *Product) (Product, error) {
	productId := -1

	for index, val := range listData {
		if val.ID == updatedUser.ID {
			productId = index
		}
	}

	if productId == -1 {
		return Product{}, errors.New("no data found")
	}

	listData[productId].Name = updatedUser.Name

	return *listData[productId], nil
}

func (model *ProductModel) Delete(deletedUser *Product) error {
	productId := -1

	for index, val := range listData {
		if val.ID == deletedUser.ID {
			productId = index
		}
	}

	if productId == -1 {
		return errors.New("no data found")
	}

	for _, val := range listData {
		if val.ID == deletedUser.ID {
			listData = append(listData[:productId], listData[productId+1:]...)
		}
	}

	return nil
}
