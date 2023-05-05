package model

import "time"

type Branch struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OwnerId     int       `json:"ownerId"`
	Title       string    `json:"title" gorm:"type:varchar(255);index"`
	Description string    `json:"description"`
	ImageBanner string    `json:"imageBanner"`
	Owner       Owner     `json:"owner" gorm:"foreignkey:OwnerId"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}
