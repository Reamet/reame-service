package model

import (
	"time"
)

type Home struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	ImageBanner string    `json:"imageBanner"`
	ButtonTitle string    `json:"buttonTitle"`
	NftIds      string    `json:"nftIds"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}
