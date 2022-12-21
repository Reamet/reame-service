package collection

import (
	"reame-service/database/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CollectionLists(c *fiber.Ctx, db *gorm.DB) error {
	offset, err := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		return err
	}

	collections := []model.Collection{}
	var count int64

	projectResult := db.Debug().Limit(limit).Offset(offset).Find(&collections)

	projectResult.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": collections,
		"total":   count,
	})
}

func CollectionById(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	collection := model.Collection{}

	db.Debug().Where("id = ?", id).First(&collection)

	return c.JSON(map[string]interface{}{
		"status": "ok",
		"result": collection,
	})
}
