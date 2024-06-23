package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/config"
	"github.com/shun198/echo-crm/models"
	"github.com/shun198/echo-crm/serializers"
	"github.com/shun198/echo-crm/services"
	"gorm.io/gorm"
)

func Login(c echo.Context, db *gorm.DB) error {
	data := new(serializers.LoginCredentials)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	user := services.GetUserByEmployeeNumber(data.EmployeeNumber, db)
	if (user == models.User{}) {
		// 処理時間から社員番号が正しいか特定できないようにする
		config.HashPassword(data.Password)
		return c.JSON(http.StatusOK, map[string]string{"msg": "社員番号またはパスワードが間違っています"})
	}
	if !config.CheckPasswordHash(user.Password, data.Password) {
		return c.JSON(http.StatusOK, map[string]string{"msg": "社員番号またはパスワードが間違っています"})
	}
	if !user.IsVerified || !user.IsActive {
		return c.JSON(http.StatusOK, map[string]string{"msg": "管理者へお問い合わせください"})
	}
	config.SetSession(user.ID, c, db)
	return c.JSON(http.StatusOK, map[string]string{})
}

func Logout(c echo.Context, db *gorm.DB) error {
	config.DeleteSession(c, db)
	return c.JSON(http.StatusOK, map[string]string{})
}
