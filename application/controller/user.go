package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/go-crm/config"
	"github.com/shun198/go-crm/model"
	"gorm.io/gorm"
)

type UserController struct {
	Env *config.Env
}

var DB *gorm.DB

func GetUsers(c echo.Context) error {
	users := []model.User{}
	DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}
