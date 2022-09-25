package service

import "01-soal-1/entity"

func NewProductService() *ProductService {
	productService := &ProductService{
		data: []entity.Product{
			{Id: 1, Name: "Benih Lele", Price: 50_000},
			{Id: 2, Name: "Pakan Lele Cap Menara", Price: 25_000},
			{Id: 3, Name: "Probiotik A", Price: 75_000},
			{Id: 4, Name: "Probiotik Nila B", Price: 10_000},
			{Id: 5, Name: "Pakan Nila", Price: 20_000},
			{Id: 6, Name: "Benih Nila Biasa", Price: 20_000},
			{Id: 7, Name: "Cupang", Price: 5_000},
			{Id: 8, Name: "Benih Nila Super", Price: 30_000},
			{Id: 9, Name: "Benih Cupang", Price: 10_000},
			{Id: 10, Name: "Probiotik B", Price: 10_000},
		},
	}
	return productService
}
