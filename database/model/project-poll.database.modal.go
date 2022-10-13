package model

import (
	"time"
)

type ProjectPoll struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement"`
	PoolId    int        `json:"pool_id"`
	Title     string     `json:"title"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
