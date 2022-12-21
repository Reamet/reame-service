package proposal

import (
	"reame-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreatePayload struct {
	PoolAddress   string    `json:"poolAddress"`
	PollId        string    `json:"pollId"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	StartVoteDate time.Time `json:"startVoteDate"`
	EndVoteDate   time.Time `json:"endVoteDate"`
}

func Create(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := CreatePayload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	proposalDatabasePayload := model.Collection{
		PoolAddress:   bodyPayload.PoolAddress,
		Title:         bodyPayload.Title,
		Description:   bodyPayload.Description,
		Status:        bodyPayload.Status,
		StartVoteDate: bodyPayload.StartVoteDate,
		EndVoteDate:   bodyPayload.EndVoteDate,
		PollId:        bodyPayload.PollId,
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
