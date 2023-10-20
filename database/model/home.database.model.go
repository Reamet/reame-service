package model

import (
	"time"
)

type Home struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string    `json:"title"`
	Subtitle      string    `json:"subtitle"`
	ImageBanner   string    `json:"image_banner"`
	ButtonTitle   string    `json:"button_title"`
	NftIds        string    `json:"nft_ids"`
	LaunchpadSlug string    `json:"launchpad_slug"`
	Launchpad     Launchpad `json:"launchpad" gorm:"foreignKey:Slug;references:LaunchpadSlug"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
