package services

import (
	"errors"
	"time"

	"github.com/shun198/echo-crm/models"
	"gorm.io/gorm"
)

func CheckInvitationToken(token string, db *gorm.DB) (models.Invitation, error) {
	var invitation models.Invitation
	result := db.Preload("User").Where("token = ?", token).Take(&invitation)
	if result.Error != nil {
		return models.Invitation{}, errors.New("存在しないトークンです")
	}

	if time.Now().After(invitation.Expiry) || invitation.Used {
		return models.Invitation{}, errors.New("有効期限切れのトークンです")
	}
	return invitation, nil
}

func CheckResetPasswordToken(token string, db *gorm.DB) (models.ResetPassword, error) {
	var reset_password models.ResetPassword
	result := db.Preload("User").Where("token = ?", token).Take(&reset_password)
	if result.Error != nil {
		return models.ResetPassword{}, errors.New("存在しないトークンです")
	}

	if time.Now().After(reset_password.Expiry) || reset_password.Used {
		return models.ResetPassword{}, errors.New("有効期限切れのトークンです")
	}
	return reset_password, nil
}

func GetUserByInvitationToken(token string, db *gorm.DB) models.User {
	var user models.User
	db.Preload("User").Where("token = ?", token)
	return user
}

func GetUserByPasswordResetToken(token string, db *gorm.DB) models.User {
	var user models.User
	db.Preload("User").Where("token = ?", token)
	return user
}
