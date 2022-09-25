package helper

import "01-soal-1/entity"

func GetLowestHighest(items []entity.Product) (entity.Product, entity.Product) {
	var lowest, highest entity.Product

	for index, item := range items {
		switch {
		case index == 0:
			highest, lowest = item, item
		case item.Price > highest.Price:
			highest = item
		case item.Price < lowest.Price:
			lowest = item
		}
	}
	return lowest, highest
}
