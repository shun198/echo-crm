package services

import (
	"errors"
	"strconv"

	"github.com/shun198/echo-crm/models"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(idParam string, db *gorm.DB) (models.User, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	result := db.First(&user, id)
	return user, result.Error
}

func GetUserByEmployeeNumber(db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Find(&user)
	return user, result.Error
}

func ToggleUserActive(user models.User, db *gorm.DB) models.User {
	db.Update("Disabled", !user.Disabled)
	return user
}

func GetRoleNumber(role string) (uint8, error) {
	switch role {
	case "管理者":
		return 1, nil
	case "一般ユーザー":
		return 2, nil
	default:
		return 99, errors.New("存在しないロールです")
	}
}
