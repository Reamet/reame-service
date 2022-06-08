package model

import (
	"time"
)

type BscProjectPayload struct {
	ReferredId											float64			`json:"id"`
	Contract                				string			`json:"contract"`
	ContractVersion									uint				`json:"contract_version"`
	OpenTime												uint				`json:"openTime"`
	CloseTime												uint				`json:"closeTime"`
	Decimals												uint				`json:"decimals"`
	State														string			`json:"state"`
	IsPrivate												bool				`json:"isPrivate"`
	Rate														string			`json:"rate"`
	TotalCountWallet								uint				`json:"totalCountWallet"`
	TotalCountUserParticipated			uint				`json:"totalCountUserParticipated"`
	TotalFundParticipated						float64			`json:"totalFundParticipated"`
	MaxSingleParticipationAllocated	uint				`json:"maxSingleParticipationAllocated"`
	MaxTotalParticipationAllocated	string			`json:"maxTotalParticipationAllocated"`
	Description											string			`json:"description"`
	Telegram												string			`json:"telegram"`
	ProjectTokenAddress							string			`json:"projectTokenAddress"`
	Logo														string			`json:"logo"`
	Medium													*string			`json:"medium"`
	Name														string			`json:"name"`
	ProjectTokenSymbol							string			`json:"projectTokenSymbol"`
	TotalSupply											string			`json:"total_supply"`
	Twitter													string			`json:"twitter"`
	Website													string			`json:"website"`
	YourAllocationVisible						bool				`json:"yourAllocationVisible"`
	Detail													*string			`json:"detail"`
	ProjectTokenContract						*string			`json:"projectTokenContract"`
	AthMultiplier										*uint				`json:"athMultiplier"`
	Symbol													string			`json:"symbol"`
	Disabled												bool				`json:"disabled"`
	Pancakeswap											string			`json:"pancakeswap"`
	Start									    	    *time.Time		`json:"start"`
	End 										        *time.Time		`json:"end"`
	Staking													string			`json:"staking"`
	Allocation											string			`json:"allocation"`
	Fcfs														string			`json:"fcfs"`
	AllDay													*bool				`json:"allDay"`
	TokenAddress										string			`json:"tokenAddress"`
}