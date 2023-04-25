package model

import (
	"time"
)

type TrendingCollection struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Ids       string    `json:"ids"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
