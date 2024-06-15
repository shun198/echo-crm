package config

import (
	"errors"
	"net/mail"
	"os"
	"time"
)

// Cookieセッションの有効期限
const sessionLength = time.Duration(2) * time.Hour

// DBセッションの有効期限
const maxSessionLength = time.Duration(10) * time.Hour

// Cookieセッションの有効期限の最大値
const maxCookieSessionLength = 60 * 60 * 10

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

var secureCookie = os.Getenv("DEBUG") == ""

var BaseDomain = os.Getenv("DOMAIN")

var cookieName = setCookieName()

func setCookieName() string {
	if secureCookie {
		return "__Host-SID"
	} else {
		return "SID"
	}
}
