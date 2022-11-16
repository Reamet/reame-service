package proposal

import (
	"bsc-scan-data-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProposalCreatePayload struct {
	PoolAddress   string    `json:"poolAddress"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	StartVoteDate time.Time `json:"startVoteDate"`
	EndVoteDate   time.Time `json:"endVoteDate"`
}

func ProposalCreate(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := ProposalCreatePayload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	proposalDatabasePayload := model.Proposal{
		PoolAddress:   bodyPayload.PoolAddress,
		Title:         bodyPayload.Title,
		Description:   bodyPayload.Description,
		Status:        bodyPayload.Status,
		StartVoteDate: bodyPayload.StartVoteDate,
		EndVoteDate:   bodyPayload.EndVoteDate,
		CreatedAt:     currentTime,
	}

	err := db.Debug().Create(&proposalDatabasePayload).Error

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
