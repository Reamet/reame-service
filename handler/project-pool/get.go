package project_pool

import (
	"bsc-scan-data-service/database/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProjectPoolList(c *fiber.Ctx, db *gorm.DB) error {
	offset, err := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))
	status := c.Query("status")

	if err != nil {
		return err
	}

	projectPools := []model.ProjectPool{}
	var count int64

	if status != "" {
		poolsResult := db.Debug().Where("status = ?", status).Preload("TierList", "deleted_at IS NULL").Limit(limit).Order("created_at desc").Offset(offset).Find(&projectPools)

		poolsResult.Debug().Offset(-1).Count(&count)

		return c.JSON(fiber.Map{
			"status":            "ok",
			"project_pool_list": projectPools,
			"amount":            count,
		})
	}

	poolsResult := db.Debug().Preload("TierList", "deleted_at IS NULL").Limit(limit).Order("created_at desc").Offset(offset).Find(&projectPools)

	poolsResult.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status":            "ok",
		"project_pool_list": projectPools,
		"amount":            count,
	})
}

func ProjectPoolById(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	projectPools := model.ProjectPool{}

	db.Debug().Where("id = ?", id).Preload("TierList", "deleted_at IS NULL").First(&projectPools)

	return c.JSON(map[string]interface{}{
		"status": "ok",
		"result": projectPools,
	})
}
