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

		databasePayload := model.BscProject {
			ID: project.ID,
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
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}