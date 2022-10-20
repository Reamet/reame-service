package pool_poll

import (
	"bsc-scan-data-service/database/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PollListByPoolId(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("poolid"))

	if err != nil {
		return err
	}

	poolPoll := []model.ProjectPollResult{}
	var count int64

	projectPollResult := db.Debug().Where("pool_id = ?", id).Order("created_at asc").Find(&poolPoll)
	projectPollResult.Debug().Offset(-1).Count(&count)

	return c.JSON(map[string]interface{}{
		"status":  "ok",
		"results": poolPoll,
		"amount":  count,
	})
}

func PollResultByPollId(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("pollid"))

	if err != nil {
		return err
	}

	pollResult := []model.ProjectPollResult{}
	var count int64

	projectPollResult := db.Debug().Where("poll_id = ?", id).Order("created_at asc").Find(&pollResult)
	projectPollResult.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": pollResult,
		"amount":  count,
	})
}
