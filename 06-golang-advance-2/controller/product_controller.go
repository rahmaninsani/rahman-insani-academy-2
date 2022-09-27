package controller

import (
	"06-golang-advance-2/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductController struct {
	Model model.ProductModel
}

func (controller *ProductController) GetProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := controller.Model.FindAll()

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get all product",
			"data":    ToProductResponses(res),
		})
	}
}

func (controller *ProductController) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.Product

		stringId := c.Param("id")
		intId, _ := strconv.Atoi(stringId)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "error when parsing data",
			})
		}
		input.ID = intId

		res, err := controller.Model.FindOne(&input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get one product",
			"data":    ToProductResponse(res),
		})
	}
}

func (controller *ProductController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newProduct model.Product

		if err := c.Bind(&newProduct); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "error when parsing data",
			})
		}

		res, err := controller.Model.Insert(&newProduct)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "error from server",
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    http.StatusCreated,
			"message": "success insert product",
			"data":    ToProductResponse(res),
		})
	}
}

func (controller *ProductController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.Product

		stringId := c.Param("id")
		intId, _ := strconv.Atoi(stringId)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "error when parsing data",
			})
		}
		input.ID = intId

		res, err := controller.Model.Update(&input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success update product",
			"data":    ToProductResponse(res),
		})
	}
}

func (controller *ProductController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.Product

		stringId := c.Param("id")
		intId, _ := strconv.Atoi(stringId)

		input.ID = intId

		err := controller.Model.Delete(&input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"code":    http.StatusNoContent,
			"message": "success delete product",
		})
	}
}
