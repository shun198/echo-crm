package config

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shun198/echo-crm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func StartDatabase() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := migrateDatabase(db); err != nil {
		return nil, err
	}
	return db, err
}

func migrateDatabase(db *gorm.DB) error {
	return db.AutoMigrate(
		models.Invitation{},
		models.ResetPassword{},
		models.User{},
		models.Session{},
	)
}

func SetSession(userId uint, c echo.Context, db *gorm.DB) error {
	key, err := tokenGenerator(64)
	if err != nil {
		return err
	}

	session := models.Session{
		UserID:    userId,
		Token:     key,
		Expiry:    time.Now().Add(sessionLength),
		MaxExpiry: time.Now().Add(maxSessionLength),
	}

	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&session)
	if result.Error != nil {
		return result.Error
	}
	WriteSessionCookie(c, key)
	return nil
}

func DeleteSession(c echo.Context, db *gorm.DB) {
	cookie, err := ReadSessionCookie(c)
	if err != nil {
		return
	}

	db.Where("token = ?", cookie).Delete(&models.Session{})
	DeleteSessionCookie(c)
}
