package model

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	PostId string `json:"postId"`
	Fish   string `json:"fish"`
	Size   string `json:"size"`
}

type Post struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Description string         `json:"description"`
	Result      Result
}
