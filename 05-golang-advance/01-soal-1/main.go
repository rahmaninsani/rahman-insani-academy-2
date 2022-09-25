package main

import (
	"01-soal-1/entity"
	"01-soal-1/helper"
	"01-soal-1/service"
	"fmt"
	"reflect"
)

func main() {
	var money uint32 = 1_000_000
	var total uint32 = money
	var cart []entity.Product

	productService := service.NewProductService()
	products, err := productService.FindAll()

	if err != nil {
		panic(err)
	}

	helper.SortAscending(products)

	for {
		for _, product := range products {
			if product.Price <= money {
				isExist := helper.SearchCartItem(cart, product.Id)

				if !isExist {
					cart = append(cart, product)
					money -= product.Price
				}
			}
		}

		if money == 0 || reflect.DeepEqual(cart, products) {
			break
		}
	}

	fmt.Println("Total produk dengan harga di bawah Rp100.000:")
	fmt.Println("Harga total", total-money)
	fmt.Println("Sisa Uang", money)
	for _, item := range cart {
		fmt.Printf(" %s - %d\n", item.Name, item.Price)
	}

	fmt.Println("=======================================")
	fmt.Println("Daftar produk dengan harga Rp10.000:")
	filteredItems := helper.FilterByPrice(products, 10_000)
	for _, item := range filteredItems {
		fmt.Printf(" %s - %d\n", item.Name, item.Price)
	}

	fmt.Println("=======================================")
	//lowest := products[0]
	//highest := products[len(products)-1]
	lowest, highest := helper.GetLowestHighest(products)

	fmt.Printf("Daftar produk termurah: %s Rp%d\n", lowest.Name, lowest.Price)
	fmt.Printf("Daftar produk termahal: %s Rp%d\n", highest.Name, highest.Price)
}
