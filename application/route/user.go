package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/go-crm/config"
)

func SetUserRoutes(env *config.Env) {
	env.Echo.GET("/api/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"msg": "pass"})
	})
}
