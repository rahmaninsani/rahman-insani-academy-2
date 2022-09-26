package main

import (
	"06-golang-advance-2/controller"
	"06-golang-advance-2/model"
	"github.com/labstack/echo/v4"
)

func main() {
	//	Echo instance
	e := echo.New()

	// Model & Controller
	productModel := model.ProductModel{}
	productController := controller.ProductController{Model: productModel}

	//	Routes
	e.GET("/products", productController.GetProducts())
	e.POST("/products", productController.Create())
	e.PUT("/products/:id", productController.Update())
	e.DELETE("/products/:id", productController.Delete())

	//	Start server
	e.Logger.Fatal(e.Start(":1323"))
}
