package config

import (
	"github.com/shun198/go-crm/model"
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
		model.Invitation{},
		model.ResetPassword{},
		model.User{},
		model.Session{},
	)
}
