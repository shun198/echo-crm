package config

import (
	"github.com/shun198/echo-crm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
