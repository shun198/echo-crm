package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/config"
	"github.com/shun198/echo-crm/emails"
	"github.com/shun198/echo-crm/models"
	"github.com/shun198/echo-crm/serializers"
	"github.com/shun198/echo-crm/services"
	"gorm.io/gorm"
)

func GetUsers(c echo.Context, db *gorm.DB, currentUser models.User) error {
	if currentUser.Role == 0 {
		users, err := services.GetAllUsers(db)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, users)
	}
	users, err := services.GetAllUsersExpectSuperuser(db)
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
	return c.JSON(http.StatusOK, map[string]bool{"is_active": toggled_user.IsActive})
}

func ChangeUserDetails(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	user, err := services.GetUserByID(id, db)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"msg": "該当するユーザが存在しません"})
	}
	data := new(serializers.ChangeUserDetails)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	services.ChangeUserDetails(user, data, db)
	return c.JSON(http.StatusOK, map[string]string{})
}

func SendInviteUserEmail(c echo.Context, db *gorm.DB) error {
	data := new(serializers.SignUp)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	// if err := c.Validate(data); err != nil {
	// 	return c.JSON(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	// }
	role, err := services.GetRoleNumber(data.Role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "存在しないロールです"})
	}
	user := services.GetUserByEmployeeNumber(data.EmployeeNumber, db)
	if (user != models.User{}) {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "登録済みのユーザです"})
	}
	user = services.GetUserByEmail(data.Email, db)
	if (user != models.User{}) {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "登録済みのユーザです"})
	}
	created_user, err := services.CreateUser(data, role, db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "ユーザの作成に失敗しました"})
	}
	invitationToken, err := config.MakeInvitationToken(&created_user, db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "招待用トークンの生成に失敗しました"})
	}
	url := config.BaseDomain + "/register/" + invitationToken.Token
	emails.SendEmail("ようこそ！", created_user.Email, url, "welcomeEmail")
	return c.JSON(http.StatusOK, map[string]string{"msg": "ユーザの招待に成功しました"})
}

func VerifyUser(c echo.Context, db *gorm.DB) error {
	data := new(serializers.ConfirmPassword)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	invitation_token, err := services.CheckInvitationToken(data.Token, db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "無効なトークンです"})
	}
	if data.NewPassword != data.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "パスワードと確認用パスワードが一致しません"})
	}
	user := invitation_token.User
	newPassword, err := config.HashPassword(data.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "パスワードのハッシュ化に失敗しました"})
	}
	db.Model(&user).Updates(models.User{IsActive: true, IsVerified: true, Password: newPassword})
	invitation_token.Used = true
	db.Save(&invitation_token)
	return c.JSON(http.StatusOK, map[string]string{})
}

func ResendInvitation(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	user, err := services.GetUserByID(id, db)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"msg": "該当するユーザが存在しません"})
	}
	url := config.BaseDomain + "/register/"
	emails.SendEmail("ようこそ！", user.Email, url, "welcomeEmail")
	return c.JSON(http.StatusOK, map[string]string{})
}

func ChangePassword(c echo.Context, db *gorm.DB) error {
	data := new(serializers.UpdateUserPassword)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	if data.NewPassword != data.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "パスワードと確認用パスワードが一致しません"})
	}
	return c.JSON(http.StatusOK, map[string]string{})
}

func SendResetPasswordEmail(c echo.Context, db *gorm.DB) error {
	data := new(serializers.SendResetPasswordEmail)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	user := services.GetUserByEmployeeNumber(data.EmployeeNumber, db)
	if (user == models.User{}) {
		return c.JSON(http.StatusBadRequest, map[string]string{})
	}
	resetPasswordToken, err := config.MakeResetPasswordToken(&user, db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{})
	}
	url := config.BaseDomain + "/reset-password/" + resetPasswordToken.Token
	emails.SendEmail("パスワード再設定", user.Email, url, "resetPassword")
	return c.JSON(http.StatusOK, map[string]string{})
}

func ResetPassword(c echo.Context, db *gorm.DB) error {
	data := new(serializers.ConfirmPassword)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	reset_password_token, err := services.CheckResetPasswordToken(data.Token, db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "無効なトークンです"})
	}
	if data.NewPassword != data.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "パスワードと確認用パスワードが一致しません"})
	}
	user := reset_password_token.User
	newPassword, err := config.HashPassword(data.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "パスワードのハッシュ化に失敗しました"})
	}
	user.Password = newPassword
	db.Save(&user)
	return c.JSON(http.StatusOK, map[string]string{})
}

func CheckInvitationToken(c echo.Context, db *gorm.DB) error {
	data := new(serializers.CheckToken)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	_, err := services.CheckInvitationToken(data.Token, db)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]bool{"check": false})
	}
	return c.JSON(http.StatusOK, map[string]bool{"check": true})
}

func CheckResetPasswordToken(c echo.Context, db *gorm.DB) error {
	data := new(serializers.CheckToken)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, serializers.ErrorResponse{Message: err.Error()})
	}
	_, err := services.CheckResetPasswordToken(data.Token, db)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]bool{"check": false})
	}
	return c.JSON(http.StatusOK, map[string]bool{"check": true})
}
