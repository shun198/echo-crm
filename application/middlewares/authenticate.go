package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/config"
	"github.com/shun198/echo-crm/serializers"
	"gorm.io/gorm"
)

var noAuthPaths = map[string]bool{
	"/api/health":                                true,
	"/api/admin/users/get_csrf_token":            true,
	"/api/admin/users/login":                     true,
	"/api/admin/users/logout":                    true,
	"/api/admin/users/verify_user":               true,
	"/api/admin/users/reset_password":            true,
	"/api/admin/user/check_invitation_token":     true,
	"/api/admin/user/check_reset_password_token": true,
}

func AuthenticateUser(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//Authがいらないパスを確認する
			if _, ok := noAuthPaths[c.Path()]; ok {
				return next(c)
			}

			token, err := config.ReadSessionCookie(c)
			if err != nil {
				return echo.ErrForbidden
			}

			user, err := config.GetUserFromSession(token, c, true, db)
			if err != nil {
				return c.JSON(http.StatusForbidden, serializers.UserInfo{Message: err.Error()})
			}

			if user != nil {
				if ok := RoutePermissions(c.Request().Method, c.Path(), user.Role); ok {
					return next(c)
				} else {
					return echo.ErrUnauthorized
				}
			} else {
				return echo.ErrForbidden
			}
		}
	}
}
