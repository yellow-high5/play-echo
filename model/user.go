package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name       string         `json:"name" validate:"required"`
	AvatarPath string         `json:"avatar"`
	Tel        string         `json:"tel" validate:"required"`
	Sex        SEX            `json:"sex" validate:"gte=0,lt=3"`
	Email      string         `json:"email" validate:"required,email"`
	Password   string         `json:"password" validate:"required"`
}
