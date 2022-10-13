package pool_poll

import (
	"bsc-scan-data-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PoolPollCreatePayload struct {
	PollId  int    `json:"poll_id"`
	PoolId  int    `json:"pool_id"`
	Address string `json:"address"`
	Value   int    `json:"value"`
}

func PollCreate(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := PoolPollCreatePayload{}

	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	poolDatabasePayload := model.ProjectPollResult{
		PollId:    bodyPayload.PollId,
		PoolId:    bodyPayload.PoolId,
		Address:   bodyPayload.Address,
		Value:     bodyPayload.Value,
		CreatedAt: currentTime,
	}

	err := db.Debug().Create(&poolDatabasePayload).Error

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
