package model

import (
	"time"
)

type Proposal struct {
	ID            int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	PoolId        int        `json:"pool_id"`
	Status        string     `json:"status"`
	StartVoteDate time.Time  `json:"start_vote_date"`
	EndVoteDate   time.Time  `json:"end_vote_date"`
	UpdatedAt     time.Time  `json:"updated_at"`
	CreatedAt     time.Time  `json:"created_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
