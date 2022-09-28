package handler

import (
	"github.com/labstack/echo/v4"
	"go-clean-arch/entity"
	"go-clean-arch/usecase"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.UserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := handler.userUseCase.CreateUser(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success insert user",
		"data":    user,
	})
}

func (handler UserHandler) GetAllUser(c echo.Context) error {
	users, err := handler.userUseCase.GetAllUser()
	if err != nil {
		return err
	}

	if users == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all user",
		"data":    users,
	})
}

func (handler UserHandler) GetUserById(c echo.Context) error {
	stringId := c.Param("id")
	intId, _ := strconv.Atoi(stringId)

	user, err := handler.userUseCase.GetUserById(intId)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get user by id",
		"data":    user,
	})
}

func (handler UserHandler) UpdateUser(c echo.Context) error {
	req := entity.UserRequest{}
	stringId := c.Param("id")
	intId, _ := strconv.Atoi(stringId)

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := handler.userUseCase.UpdateUser(req, intId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success update user",
		"data":    user,
	})
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	stringId := c.Param("id")
	intId, _ := strconv.Atoi(stringId)

	err := handler.userUseCase.DeleteUser(intId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success delete user",
	})
}
