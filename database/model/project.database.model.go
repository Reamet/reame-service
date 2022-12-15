package model

import (
	"time"
)

type Project struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Information string     `json:"information"`
	Source      string     `json:"source"`
	Logo        string     `json:"logo"`
	Website     string     `json:"website"`
	Status      string     `json:"status"`
	Telegram    *string    `json:"telegram"`
	Twitter     *string    `json:"twitter"`
	Discord     *string    `json:"discord"`
	Email       *string    `json:"email"`
	Facebook    *string    `json:"facebook"`
	Instagram   *string    `json:"instagram"`
	ReferredId  *float64   `json:"referred_id"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
