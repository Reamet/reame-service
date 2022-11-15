package proposal

import (
	"bsc-scan-data-service/database/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProposalLists(c *fiber.Ctx, db *gorm.DB) error {
	offset, err := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		return err
	}

	results := []model.Proposal{}
	var count int64

	proposalResult := db.Debug().Limit(limit).Offset(offset).Find(&results)

	proposalResult.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": results,
		"amount":  count,
	})
}

func ProposalById(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	result := model.Proposal{}

	db.Debug().Where("id = ?", id).First(&result)

	return c.JSON(fiber.Map{
		"status": "ok",
		"result": result,
	})
}
