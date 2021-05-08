package model

import (
	"time"

	"gorm.io/gorm"
)

type Reserve struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	UserId     string         `json:"userId" validate:"required"`
	PlanId     string         `json:"planId" validate:"required"`
	Number     uint           `json:"number" validate:"required"`
	IsCanceled bool           `json:"isCanceled" validate:"required"`
	IsExecuted bool           `json:"isExecuted" validate:"required"`
}
