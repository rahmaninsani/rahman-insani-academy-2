package helper

import "01-soal-1/entity"

func FilterByPrice(items []entity.Product, price uint32) []entity.Product {
	var filteredItems []entity.Product

	for _, item := range items {
		if item.Price == price {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
