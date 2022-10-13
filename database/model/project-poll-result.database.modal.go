package model

import (
	"time"
)

type ProjectPollResult struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	PollId    int       `json:"poll_id"`
	PoolId    int       `json:"pool_id"`
	Address   string    `json:"address"`
	Value     int       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}
