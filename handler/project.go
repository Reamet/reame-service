package handler

import (
	"bsc-scan-data-service/database/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProjectHandler struct {
	DB *gorm.DB
}

func (ph *ProjectHandler) Init(db *gorm.DB) {
	ph.DB = db
}

type ProjectPayload struct {
	ProjectList		[]model.BscProjectPayload		`json:"projectList"`
}

func (ph *ProjectHandler) ProjectCreate(c *fiber.Ctx) error {

	bodyPayload := ProjectPayload{}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}



	for _,project := range bodyPayload.ProjectList {

		projectResult := model.BscProject{}

		result := ph.DB.Debug().Where("referred_id = ?", project.ReferredId).First(&projectResult)
		
		if result.RowsAffected == 0 {
			databasePayload := model.BscProject {
				ReferredId: project.ReferredId,
				Contract: project.Contract,
				ContractVersion: project.ContractVersion,
				OpenTime: project.OpenTime,
				CloseTime: project.CloseTime,
				Decimals: project.Decimals,
				State: project.State,
				IsPrivate: project.IsPrivate,
				Rate: project.Rate,
				TotalCountWallet: project.TotalCountWallet,
				TotalCountUserParticipated: project.TotalCountUserParticipated,
				TotalFundParticipated: project.TotalFundParticipated,
				MaxSingleParticipationAllocated: project.MaxSingleParticipationAllocated,
				MaxTotalParticipationAllocated: project.MaxTotalParticipationAllocated,
				Description: project.Description,
				Telegram: project.Telegram,
				ProjectTokenAddress: project.ProjectTokenAddress,
				Logo: project.Logo,
				Medium: project.Medium,
				Name: project.Name,
				ProjectTokenSymbol: project.ProjectTokenSymbol,
				TotalSupply: project.TotalSupply,
				Twitter: project.Twitter,
				Website: project.Website,
				YourAllocationVisible: project.YourAllocationVisible,
				Detail: project.Detail,
				ProjectTokenContract: project.ProjectTokenContract,
				AthMultiplier: project.AthMultiplier,
				Symbol: project.Symbol,
				Disabled: project.Disabled,
				Pancakeswap: project.Pancakeswap,
				Start: project.Start,
				End: project.End,
				Staking: project.Staking,
				Allocation: project.Allocation,
				Fcfs: project.Fcfs,
				AllDay: project.AllDay,
				TokenAddress: project.TokenAddress,
			}
	
			err := ph.DB.Debug().Create(&databasePayload).Error
			if err != nil {
				return err
			}
		} else {
			databasePayload := map[string]interface{}  {
				"referred_id": project.ReferredId,
				"contract": project.Contract,
				"contract_version": project.ContractVersion,
				"open_time": project.OpenTime,
				"close_time": project.CloseTime,
				"decimals": project.Decimals,
				"state": project.State,
				"is_private": project.IsPrivate,
				"rate": project.Rate,
				"total_count_wallet": project.TotalCountWallet,
				"total_count_user_participated": project.TotalCountUserParticipated,
				"total_fund_participated": project.TotalFundParticipated,
				"max_single_participation_allocated": project.MaxSingleParticipationAllocated,
				"max_total_participation_allocated": project.MaxTotalParticipationAllocated,
				"description": project.Description,
				"telegram": project.Telegram,
				"project_token_address": project.ProjectTokenAddress,
				"logo": project.Logo,
				"medium": project.Medium,
				"name": project.Name,
				"project_token_symbol": project.ProjectTokenSymbol,
				"total_supply": project.TotalSupply,
				"twitter": project.Twitter,
				"website": project.Website,
				"your_allocation_visible": project.YourAllocationVisible,
				"detail": project.Detail,
				"project_token_contract": project.ProjectTokenContract,
				"ath_multiplier": project.AthMultiplier,
				"symbol": project.Symbol,
				"disabled": project.Disabled,
				"pancakeswap": project.Pancakeswap,
				"start": project.Start,
				"end": project.End,
				"staking": project.Staking,
				"allocation": project.Allocation,
				"fcfs": project.Fcfs,
				"all_day": project.AllDay,
				"token_address": project.TokenAddress,
			}

			err := result.Debug().Updates(&databasePayload).Error
			if err != nil {
				return err
			}
		}

	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}