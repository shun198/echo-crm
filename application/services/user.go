package services

import (
	"errors"
	"strconv"

	"github.com/shun198/echo-crm/models"
	"github.com/shun198/echo-crm/serializers"
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

func GetUserByEmployeeNumber(employee_number string, db *gorm.DB) models.User {
	var user models.User
	result := db.Where("employee_number = ?", employee_number).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return models.User{}
	}
	return user
}

func GetUserByEmail(email string, db *gorm.DB) models.User {
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return models.User{}
	}
	return user
}

func ToggleUserActive(user models.User, db *gorm.DB) models.User {
	db.Update("Disabled", !user.IsActive)
	return user
}

func CreateUser(data *serializers.SignUp, role uint8, db *gorm.DB) (models.User, error) {
	user := models.User{Name: data.Name, Email: data.Email, EmployeeNumber: data.EmployeeNumber, Role: role}
	result := db.Create(&user)
	return user, result.Error
}

func GetRoleNumber(role string) (uint8, error) {
	switch role {
	case "管理者":
		return 1, nil
	case "一般":
		return 2, nil
	default:
		return 99, errors.New("存在しないロールです")
	}
}
