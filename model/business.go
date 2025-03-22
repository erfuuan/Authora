package model

import (
	"time"
)

type Business struct {
	ID         unit      `gorm:"primary_key;autoIncrement"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Token      string    `gorm:"type:varchar(255);not null;unique"`
	Created_at time.Time `gorm:"default:current_timestamp"`
	Updated_at time.Time `gorm:"default:null"`
	Deleted_at time.Time `gorm:"default:null"`
}
