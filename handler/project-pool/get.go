package project_pool

import (
	"bsc-scan-data-service/database/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProjectPoolList(c *fiber.Ctx, db *gorm.DB) error {
	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		return err
	}

	projectPools := []model.ProjectPool{}

	poolsResult := db.Debug().Preload("TierList", "deleted_at IS NULL").Limit(10).Order("created_at desc").Offset(offset).Find(&projectPools)

	var count int64

	poolsResult.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status":            "ok",
		"project_pool_list": projectPools,
		"amount":            count,
	})
}
