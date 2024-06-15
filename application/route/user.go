package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/controllers"
	"gorm.io/gorm"
)

func SetUserRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/api/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"msg": "pass"})
	})
	e.GET("/api/admin/users", func(c echo.Context) error {
		return controllers.GetUsers(c, db)
	})
	e.POST("/api/admin/users/:id/toggle_user_active", func(c echo.Context) error {
		return controllers.ToggleUserActive(c, db)
	})
}
