package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	HostId      string         `json:"hostId" validate:"required"`
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description"`
	Price       uint           `json:"price" validate:"required"`
	Type        PLANTYPE       `json:"type" validate:"required"`
	TargetFish  string         `json:"targetFish"`
	Capasity    int32          `json:"capasity" validate:"required"`
	MeetingAt   time.Time      `json:"meetingAt" validate:"required"`
	DepartureAt time.Time      `json:"departureAt" validate:"required"`
	ReturnAt    time.Time      `json:"returnAt" validate:"required"`
	Duration    time.Duration  `json:"duration" validate:"required"`
}
