package services

import (
	"github.com/shun198/echo-crm/models"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}
