package route

import (
	"net/http"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/config"
	"github.com/shun198/echo-crm/controllers"
	"gorm.io/gorm"
)

func SetUserRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/api/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"msg": "pass"})
	})
	e.GET("/api/admin/users/get_csrf_token", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"csrf_token": c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)})
	})
	e.POST("/api/admin/users/login", func(c echo.Context) error {
		return controllers.Login(c, db)
	})
	e.POST("/api/admin/users/logout", func(c echo.Context) error {
		return controllers.Logout(c, db)
	})
	e.GET("/api/admin/users", func(c echo.Context) error {
		currentUser := config.GetUserFromCookie(c, db)
		return controllers.GetUsers(c, db, *currentUser)
	})
	e.POST("/api/admin/users/:id/toggle_user_active", func(c echo.Context) error {
		return controllers.ToggleUserActive(c, db)
	})
	e.PATCH("/api/admin/users/:id/change_user_details", func(c echo.Context) error {
		return controllers.ToggleUserActive(c, db)
	})
	e.POST("/api/admin/users/send_invite_user_email", func(c echo.Context) error {
		return controllers.SendInviteUserEmail(c, db)
	})
	e.POST("/api/admin/users/verify_user", func(c echo.Context) error {
		return controllers.VerifyUser(c, db)
	})
	e.POST("/api/admin/users/:id/resend_invitation", func(c echo.Context) error {
		return controllers.ResendInvitation(c, db)
	})
	e.POST("/api/admin/users/change_password", func(c echo.Context) error {
		return controllers.ChangePassword(c, db)
	})
	e.POST("/api/admin/users/send_reset_password_email", func(c echo.Context) error {
		return controllers.SendResetPasswordEmail(c, db)
	})
	e.POST("/api/admin/users/reset_password", func(c echo.Context) error {
		return controllers.ResetPassword(c, db)
	})
	e.POST("/api/admin/users/check_invitation_token", func(c echo.Context) error {
		return controllers.CheckInvitationToken(c, db)
	})
	e.POST("/api/admin/users/check_reset_password_token", func(c echo.Context) error {
		return controllers.CheckResetPasswordToken(c, db)
	})
}

type PermissionSet map[string][]int

func GetUserPermission() map[string]PermissionSet {
	return map[string]PermissionSet{
		"/api/admin/users":                              {"GET": config.AdminPermission()},
		"/api/admin/user/toggle_user_active":            {"POST": config.AdminPermission()},
		"/api/admin/user/:id/change_user_details":       {"PATCH": config.AdminPermission()},
		"/api/admin/user/:id/resend_invitation":         {"POST": config.AdminPermission()},
		"/api/admin/user/:id/send_invite_user_email":    {"POST": config.AdminPermission()},
		"/api/admin/user/:id/send_reset_password_email": {"POST": config.AdminPermission()},
	}
}
