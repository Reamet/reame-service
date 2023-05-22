package model

import (
	"time"
)

type Launchpad struct {
	ID               string    `json:"id" gorm:"type:varchar(42);primary_key"`
	Title            string    `json:"title" gorm:"type:varchar(255);index"`
	Description      string    `json:"description"`
	ImageBanner      string    `json:"imageBanner"`
	ImageFeature     string    `json:"imageFeature"`
	ImageAvatar      string    `json:"imageAvatar"`
	ImageSlider      string    `json:"imageSlider"`
	ChainName        string    `json:"chainName"`
	LaunchpadAddress string    `json:"launchpadAddress"`
	Hot              *bool     `json:"hot"`
	Slug             string    `json:"slug" gorm:"unique;type:varchar(255)"`
	Instragram       string    `json:"instragram" gorm:"type:varchar(255)"`
	Twitter          string    `json:"twitter" gorm:"type:varchar(255)"`
	Facebook         string    `json:"facebook" gorm:"type:varchar(255)"`
	Telegram         string    `json:"telegram" gorm:"type:varchar(255)"`
	Discord          string    `json:"discord" gorm:"type:varchar(255)"`
	CreatedBy        string    `json:"createdBy" gorm:"type:varchar(42)"`
	UpdateBy         string    `json:"updateBy" gorm:"type:varchar(42)"`
	Active           *bool     `json:"active"`
	TermAndCondition string    `json:"termAndCondition"`
	Status           string    `json:"status"`
	SaleStatus       string    `json:"saleStatus"`
	SaleType         string    `json:"saleType"`
	StartDate        time.Time `json:"startDate,omitempty"`
	EndDate          time.Time `json:"endDate,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
}
