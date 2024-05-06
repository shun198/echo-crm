package config

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"github.com/shun198/go-crm/model"
)

// パスワード再設定用トークンを作るための関数
func MakeResetPasswordToken(user *model.User) (*model.ResetPassword, error) {
	token, err := tokenGenerator(32)
	if err != nil {
		return nil, err
	}

	expiryTime := time.Now().Add(resetPasswordExpiry)
	resetPassword := model.ResetPassword{
		User:   *user,
		UserID: user.ID,
		Token:  token,
		Expiry: expiryTime,
		Used:   false,
	}
	return &resetPassword, nil
}

// 招待用トークンを作るための関数
func MakeInvitationToken(user *model.User) (*model.Invitation, error) {
	token, err := tokenGenerator(32)
	if err != nil {
		return nil, err
	}
	expiryTime := time.Now().Add(userInviteExpiry)
	invitation := model.Invitation{
		User:   *user,
		UserID: user.ID,
		Token:  token,
		Expiry: expiryTime,
		Used:   false,
	}
	return &invitation, nil
}

func RandomPassword() (string, error) {
	token, err := tokenGenerator(32)
	if err != nil {
		return "", err
	}
	return token, nil
}

// トークン作成用
func tokenGenerator(length int) (string, error) {
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
