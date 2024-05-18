package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/go-crm/config"
	"github.com/shun198/go-crm/controller"
)

func SetUserRoutes(env *config.Env) {
	uc := controller.UserController{Env: env}

	env.Echo.GET("/api/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"msg": "pass"})
	})

	env.Echo.GET("/api/user", func(c echo.Context) error {
		users := uc.GetUsers()
		return c.JSON(http.StatusOK, users)
	})
}
