package main

import (
	"fmt"
	"time"

	"github.com/shun198/echo-crm/config"
	"github.com/shun198/echo-crm/models"
)

func main() {
	var Env config.Env
	password, _ := config.HashPassword("test")
	admin_user := models.User{
		EmployeeNumber: "00000001",
		Email:          "test01@example.com",
		Name:           "テストユーザ01",
		Password:       password,
		Verified:       true,
		Role:           0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	// Userデータをデータベースに保存
	if err := Env.DB.Create(&admin_user).Error; err != nil {
		fmt.Println("ユーザの新規作成に失敗しました:", err)
	}
	general_user := models.User{
		EmployeeNumber: "00000002",
		Email:          "test02@example.com",
		Name:           "テストユーザ02",
		Password:       password,
		Verified:       true,
		Role:           1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	// Userデータをデータベースに保存
	if err := Env.DB.Create(&general_user).Error; err != nil {
		fmt.Println("ユーザの新規作成に失敗しました:", err)
	}
	fmt.Println("テストデータの作成に成功しました")
}
