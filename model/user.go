package model

import (
	"time"
)

type User struct {
	ID         int       `gorm:"primary_key;autoIncrement"`
	UserId     string    `gorm:"type:varchar(255);not null"`
	ChatId     int64     `gorm:"not null"`
	Created_at time.Time `gorm:"default:current_timestamp"`
	Updated_at time.Time `gorm:"default:null"`
	Deleted_at time.Time `gorm:"default:null"`
}
