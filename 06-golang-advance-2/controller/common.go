package controller

import (
	"06-golang-advance-2/model"
)

type ProductResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToProductResponse(model model.Product) ProductResponse {
	return ProductResponse{
		ID:   model.ID,
		Name: model.Name,
	}
}

func ToProductResponses(model []*model.Product) []ProductResponse {
	var res []ProductResponse

	for _, val := range model {
		res = append(res, ProductResponse{ID: val.ID, Name: val.Name})
	}

	return res
}
