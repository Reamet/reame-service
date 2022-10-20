package pool_poll

import (
	"bsc-scan-data-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PoolPollCreatePayload struct {
	PollId  int    `json:"pollId"`
	PoolId  int    `json:"poolId"`
	Address string `json:"address"`
	Value   int    `json:"value"`
}

func PollCreate(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := PoolPollCreatePayload{}
	var count int64
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	pollResult := model.ProjectPollResult{}
	projectPollResult := db.Debug().Where("address = ?", bodyPayload.Address).First(&pollResult)
	projectPollResult.Debug().Offset(-1).Count(&count)

	if count > 0 {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "duplicate address",
		})
	}

	poolDatabasePayload := model.ProjectPollResult{
		PollId:    int(bodyPayload.PollId),
		PoolId:    int(bodyPayload.PoolId),
		Address:   bodyPayload.Address,
		Value:     int(bodyPayload.Value),
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
