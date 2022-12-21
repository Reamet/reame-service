package model

import (
	"time"
)

type Mint struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	MintImage   string     `json:"mint_image"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Collection  string     `json:"collection"`
	Address     string     `json:"address"`
	Royalties   string     `json:"royalties"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
