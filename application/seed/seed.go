package main

import (
	"fmt"
	"time"

	"github.com/shun198/echo-crm/config"
	"github.com/shun198/echo-crm/models"
	"gorm.io/gorm"
)

func main() {
	var db gorm.DB
	password, _ := config.HashPassword("test")
	admin_user := models.User{
		EmployeeNumber: "00000001",
		Email:          "test01@example.com",
		Name:           "テストユーザ01",
		Password:       password,
		IsVerified:     true,
		Role:           0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	// Userデータをデータベースに保存
	if err := db.Create(&admin_user).Error; err != nil {
		fmt.Println("ユーザの新規作成に失敗しました:", err)
	}
	general_user := models.User{
		EmployeeNumber: "00000002",
		Email:          "test02@example.com",
		Name:           "テストユーザ02",
		Password:       password,
		IsVerified:     true,
		Role:           1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	// Userデータをデータベースに保存
	if err := db.Create(&general_user).Error; err != nil {
		fmt.Println("ユーザの新規作成に失敗しました:", err)
	}
	fmt.Println("テストデータの作成に成功しました")
}
