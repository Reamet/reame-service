package model

import (
	"time"

	"github.com/lib/pq"
)

type ProjectPool struct {
	ID                        int           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title                     string        `json:"title"`
	SubTitle                  string        `json:"sub_title"`
	Description               string        `json:"description"`
	Source                    string        `json:"source"`
	StartDate                 time.Time     `json:"start_date"`
	EndDate                   time.Time     `json:"end_date"`
	ProjectList               pq.Int64Array `gorm:"type:integer[]" json:"project_list"`
	Term                      string        `json:"term"`
	InvestmentPeriod          int           `json:"investment_period"`
	WithdrawalDate            time.Time     `json:"withdrawal_date"`
	GoalRaised                int           `json:"goal_raised"`
	GoalAllocation            int           `json:"goal_allocation"`
	BasicInvestmentSuggestion int           `json:"basic_investment_suggestion"`
	DepositFee                int           `json:"deposit_fee"`
	Ido                       int           `json:"ido"`
	Stake                     int           `json:"stake"`
	Status                    string        `json:"status"`
	TierList                  []ProjectTier `gorm:"foreignKey:PoolId" json:"tier_list"`
	UpdatedAt                 time.Time     `json:"updated_at"`
	CreatedAt                 time.Time     `json:"created_at"`
	DeletedAt                 *time.Time    `json:"deleted_at"`
}

type ProjectPoolResponse struct {
	ID                        int           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title                     string        `json:"title"`
	SubTitle                  string        `json:"sub_title"`
	Description               string        `json:"description"`
	Source                    string        `json:"source"`
	StartDate                 time.Time     `json:"start_date"`
	EndDate                   time.Time     `json:"end_date"`
	ProjectList               []Project     `json:"project_list"`
	Term                      string        `json:"term"`
	InvestmentPeriod          int           `json:"investment_period"`
	WithdrawalDate            time.Time     `json:"withdrawal_date"`
	GoalRaised                int           `json:"goal_raised"`
	GoalAllocation            int           `json:"goal_allocation"`
	BasicInvestmentSuggestion int           `json:"basic_investment_suggestion"`
	DepositFee                int           `json:"deposit_fee"`
	Ido                       int           `json:"ido"`
	Stake                     int           `json:"stake"`
	Status                    string        `json:"status"`
	TierList                  []ProjectTier `gorm:"foreignKey:PoolId" json:"tier_list"`
	UpdatedAt                 time.Time     `json:"updated_at"`
	CreatedAt                 time.Time     `json:"created_at"`
	DeletedAt                 *time.Time    `json:"deleted_at"`
}
