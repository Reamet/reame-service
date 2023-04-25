package model

import (
	"time"
)

type FeaturedCollection struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Ids       string    `json:"ids"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
