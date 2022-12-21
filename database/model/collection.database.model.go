package model

import (
	"time"
)

type Collection struct {
	ID                     int        `json:"id" gorm:"primaryKey;autoIncrement"`
	CollectionProfileImage string     `json:"collection_profile_image"`
	CollectionCoverImage   string     `json:"collection_cover_image"`
	Name                   string     `json:"name"`
	Description            string     `json:"description"`
	ShortUrl               string     `json:"short_url"`
	Category               string     `json:"category"`
	Address                string     `json:"address"`
	Website                string     `json:"website"`
	Facebook               string     `json:"facebook"`
	Twitter                string     `json:"twitter"`
	Discord                string     `json:"discord"`
	Telegram               string     `json:"telegram"`
	Medium                 string     `json:"medium"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              *time.Time `json:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at"`
}
