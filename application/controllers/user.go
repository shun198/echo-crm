package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/models"
	"github.com/shun198/echo-crm/serializers"
	"github.com/shun198/echo-crm/services"
	"gorm.io/gorm"
)

func GetUsers(c echo.Context, db *gorm.DB) error {
	users, err := services.GetAllUsers(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func ToggleUserActive(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	user, err := services.GetUserByID(id, db)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"msg": "該当するユーザが存在しません"})
	}
	toggled_user := services.ToggleUserActive(user, db)
	return c.JSON(http.StatusOK, map[string]bool{"disabled": toggled_user.Active})
}

func SendInviteUserEmail(c echo.Context, db *gorm.DB) error {
	data := new(serializers.SignUp)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	// if err := c.Validate(data); err != nil {
	// 	return c.JSON(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	// }
	_, err := services.GetRoleNumber(data.Role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "存在しないロールです"})
	}
	user, _ := services.GetUserByEmployeeNumber(data.EmployeeNumber, db)
	if (user == models.User{}) {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "登録済みのユーザです"})
	}
	services.CreateUser(data, db)
	return c.JSON(http.StatusOK, map[string]string{"msg": "ユーザの招待に成功しました"})
}

func ResendInvitation(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	user, err := services.GetUserByID(id, db)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"msg": "該当するユーザが存在しません"})
	}
	fmt.Print(user)
	return c.JSON(http.StatusOK, map[string]string{})
}
