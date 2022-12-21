package mint

import (
	"reame-service/database/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MintLists(c *fiber.Ctx, db *gorm.DB) error {
	offset, err := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))
	address := c.Query("address")
	category := c.Query("category")

	if err != nil {
		return err
	}

	mints := []model.Mint{}
	var count int64

	if address != "" || category == "" {
		result := db.Debug().Where("address = ?", address).Limit(limit).Offset(offset).Find(&mints)

		result.Debug().Offset(-1).Count(&count)

		return c.JSON(fiber.Map{
			"status":  "ok",
			"results": mints,
			"total":   count,
		})
	}

	if address == "" || category != "" {
		result := db.Debug().Where("category = ?", category).Limit(limit).Offset(offset).Find(&mints)

		result.Debug().Offset(-1).Count(&count)

		return c.JSON(fiber.Map{
			"status":  "ok",
			"results": mints,
			"total":   count,
		})
	}

	if address != "" && category != "" {
		result := db.Debug().Where("address = ? AND category = ?", address, category).Limit(limit).Offset(offset).Find(&mints)

		result.Debug().Offset(-1).Count(&count)

		return c.JSON(fiber.Map{
			"status":  "ok",
			"results": mints,
			"total":   count,
		})
	}

	result := db.Debug().Limit(limit).Offset(offset).Find(&mints)

	result.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": mints,
		"total":   count,
	})
}

func GetMintById(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	mint := model.Mint{}

	db.Debug().Where("id = ?", id).First(&mint)

	return c.JSON(map[string]interface{}{
		"status": "ok",
		"result": mint,
	})
}
