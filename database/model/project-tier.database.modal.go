package model

import (
	"time"
)

type ProjectTier struct {
	ID                        int           `json:"id" gorm:"primaryKey;autoIncrement"`
	PoolId										int						`json:"pool_id"`
	Tier											int						`json:"tier"`
	TokenAmount								int						`json:"token_amount"`
	UpdatedAt                 time.Time     `json:"updated_at"`
	CreatedAt                 time.Time     `json:"created_at"`
	DeletedAt                 *time.Time    `json:"deleted_at"`
}
