package model

import (
	"time"
)

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	EmployeeNumber string    `gorm:"uniqueIndex;not null;size:8" json:"employeeNumber"`
	Email          string    `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Name           string    `gorm:"not null;size:255" json:"name"`
	Password       string    `gorm:"not null;size:255" json:"-"`
	CreatedAt      time.Time `gorm:"- autoCreateTime" json:"-"`
	Verified       bool      `gorm:"not null" json:"verified"`
	Role           uint8     `gorm:"not null" json:"role"`
	Disabled       bool      `gorm:"not null;default:false"`
	UpdatedAt      time.Time `gorm:"- autoUpdateTime" json:"-"`
	UpdatedByID    *uint     `gorm:"" json:"-"`
	UpdatedBy      *User     `gorm:"not null" json:"-"`
	CreatedByID    *uint     `gorm:"" json:"-"`
	CreatedBy      *User     `gorm:"not null" json:"-"`
}
