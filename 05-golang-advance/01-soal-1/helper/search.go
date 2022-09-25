package helper

import "01-soal-1/entity"

func SearchCartItem(items []entity.Product, productId uint64) bool {
	for _, item := range items {
		if item.Id == productId {
			return true
		}
	}
	return false
}
