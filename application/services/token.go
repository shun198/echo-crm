package services

import (
	"time"

	"github.com/shun198/echo-crm/models"
	"gorm.io/gorm"
)

func CheckInvitationToken(token string, db *gorm.DB) bool {
	var invitation *models.Invitation
	result := db.Preload("User").Where("token = ?", token).Take(&invitation)
	if result.Error != nil {
		return false
	}

	if time.Now().After(invitation.Expiry) || invitation.Used {
		return false
	}
	return true
}

func CheckResetPasswordToken(token string, db *gorm.DB) bool {
	var reset_password *models.ResetPassword
	result := db.Preload("User").Where("token = ?", token).Take(&reset_password)
	if result.Error != nil {
		return false
	}

	if time.Now().After(reset_password.Expiry) || reset_password.Used {
		return false
	}
	return true
}
