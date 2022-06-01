package model

import (
	"time"
)

type BscProject struct {
	ID                							float64			`json:"id"`
	Contract                				string			`json:"contract"`
	ContractVersion									uint				`json:"contract_version"`
	OpenTime												uint				`json:"open_time"`
	CloseTime												uint				`json:"close_time"`
	Decimals												uint				`json:"decimals"`
	State														string			`json:"state"`
	IsPrivate												bool				`json:"is_private"`
	Rate														string			`json:"rate"`
	TotalCountWallet								uint				`json:"total_countWallet"`
	TotalCountUserParticipated			uint				`json:"total_count_user_participated"`
	TotalFundParticipated						float64			`json:"total_fund_participated"`
	MaxSingleParticipationAllocated	uint				`json:"max_single_participation_allocated"`
	MaxTotalParticipationAllocated	string			`json:"max_total_participation_allocated"`
	Description											string			`json:"description"`
	Telegram												string			`json:"telegram"`
	ProjectTokenAddress							string			`json:"project_token_address"`
	Logo														string			`json:"logo"`
	Medium													string			`json:"medium"`
	Name														string			`json:"name"`
	ProjectTokenSymbol							string			`json:"project_token_symbol"`
	TotalSupply											string			`json:"total_supply"`
	Twitter													string			`json:"twitter"`
	Website													string			`json:"website"`
	YourAllocationVisible						bool				`json:"your_allocation_visible"`
	Detail													string			`json:"detail"`
	ProjectTokenContract						string			`json:"project_token_contract"`
	AthMultiplier										uint				`json:"ath_multiplier"`
	Symbol													string			`json:"symbol"`
	Disabled												bool				`json:"disabled"`
	Pancakeswap											string			`json:"pancakeswap"`
	Start									    	    time.Time		`json:"start"`
	End 										        time.Time		`json:"end"`
	Staking													string			`json:"staking"`
	Allocation											string			`json:"allocation"`
	Fcfs														string			`json:"fcfs"`
	AllDay													bool				`json:"all_day"`
	TokenAddress										string			`json:"token_address"`
}