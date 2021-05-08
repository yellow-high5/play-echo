package model

import (
	"time"

	"gorm.io/gorm"
)

type Host struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name             string         `json:"name" validate:"required"`
	Representative   string         `json:"representative" validate:"required"`
	AvatarPath       string         `json:"avatar"`
	UrlPath          string         `json:"url"`
	Tel              string         `json:"tel" validate:"required"`
	Address          string         `json:"address" validate:"required"`
	Plans            []Plan         `json:"plans" gorm:"foreignKey:HostId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Intro            string         `json:"intro"`
	HasParking       bool           `json:"hasParking"`
	HasCreditPayment bool           `json:"hasCreditPayment"`
}
