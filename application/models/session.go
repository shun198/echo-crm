package models

import (
	"time"
)

type Session struct {
	ID        uint      `gorm:"primaryKey"`
	User      User      `gorm:"not null;constraint:OnDelete:CASCADE"`
	UserID    uint      `gorm:"not null;unique"`
	Token     string    `gorm:"unique;not null;size:255"`
	CreatedAt time.Time `gorm:"- autoCreateTime"`
	Expiry    time.Time `gorm:"not null"`
	MaxExpiry time.Time `gorm:"not null"`
}
