package project

import (
	"bsc-scan-data-service/database/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProjectLists(c *fiber.Ctx, db *gorm.DB) error {
	offset, err := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))
	status := c.Query("status")

	if err != nil {
		return err
	}

	project := []model.Project{}
	var count int64

	if status != "" {
		projectResult := db.Debug().Where("status = ?", status).Limit(limit).Offset(offset).Find(&project)

		projectResult.Debug().Offset(-1).Count(&count)

		return c.JSON(fiber.Map{
			"status":        "ok",
			"project_lists": project,
			"amount":        count,
		})
	}

	projectResult := db.Debug().Limit(limit).Offset(offset).Find(&project)

	projectResult.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status":        "ok",
		"project_lists": project,
		"amount":        count,
	})
}

func ProjectById(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	project := model.Project{}

	db.Debug().Where("id = ?", id).First(&project)

	return c.JSON(map[string]interface{}{
		"status": "ok",
		"result": project,
	})
}
