package helper

import "01-soal-1/entity"

func SortAscending(items []entity.Product) []entity.Product {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items); j++ {
			if items[i].Price < items[j].Price {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	return items
}
