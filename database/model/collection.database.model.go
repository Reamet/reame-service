package model

import (
	"time"
)

type Collection struct {
	ID                string    `json:"id" gorm:"type:varchar(42);primary_key"`
	CollectionIdChain *string   `json:"collection_id_chain"`
	Title             string    `json:"title" gorm:"type:varchar(255);index"`
	Description       string    `json:"description"`
	ImageBanner       string    `json:"imageBanner"`
	ImageFeature      *string   `json:"image_feature"`
	ImageAvatar       string    `json:"image_avatar"`
	TokenType         *string   `json:"token_type"`
	Hot               *bool     `json:"hot"`
	OwnerId           int       `json:"owner_id"`
	BranchId          *int      `json:"branch_id"`
	Slug              string    `json:"slug" gorm:"unique;type:varchar(255)"`
	Instragram        string    `json:"instragram" gorm:"type:varchar(255)"`
	Twitter           string    `json:"twitter" gorm:"type:varchar(255)"`
	Facebook          string    `json:"facebook" gorm:"type:varchar(255)"`
	Telegram          string    `json:"telegram" gorm:"type:varchar(255)"`
	Discord           string    `json:"discord" gorm:"type:varchar(255)"`
	CreatedBy         string    `json:"created_by" gorm:"type:varchar(42)"`
	UpdateBy          string    `json:"update_by" gorm:"type:varchar(42)"`
	Active            string    `json:"active"`
	TermAndCondition  string    `json:"term_and_condition"`
	Status            string    `json:"status"`
	Owner             Owner     `json:"owner" gorm:"foreignkey:OwnerId"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}
