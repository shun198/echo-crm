package config

import (
	"errors"
	"net/mail"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// EchoとGormのポインタを格納
type Env struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

// パスワード再設定用トークンの有効期限
const resetPasswordExpiry = time.Duration(30) * time.Minute

// ユーザ招待用トークンの有効期限
const userInviteExpiry = time.Duration(24) * time.Hour

func ValidateEmail(email string) error {
	if email == "" {
		return nil
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("正しい形式のメールアドレスを入力してください")
	} else {
		return nil
	}
}
