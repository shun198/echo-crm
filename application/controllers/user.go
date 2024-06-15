package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
