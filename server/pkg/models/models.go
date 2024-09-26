package models

import (
	"time"
)

type Users struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"size:100" validate:"required" json:"name"`
	Email         string    `gorm:"uniqueIndex" validate:"required,email" json:"email" binding:"required"`
	Surname       string    `gorm:"size:100" validate:"required" json:"surname"`
	Password      string    `gorm:"size:255" validate:"required" json:"password" binding:"required"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
	ProfileImgUrl string    `json:"profileImg"`
}
