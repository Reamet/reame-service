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
	address := c.Query("address")

	if err != nil {
		return err
	}

	collections := []model.Collection{}
	var count int64

	if address != "" {
		results := db.Debug().Where("address = ?", address).Limit(limit).Offset(offset).Find(&collections)

		results.Debug().Offset(-1).Count(&count)

		return c.JSON(fiber.Map{
			"status":  "ok",
			"results": collections,
			"total":   count,
		})
	}

	results := db.Debug().Limit(limit).Offset(offset).Find(&collections)

	results.Debug().Offset(-1).Count(&count)
	if results.Error == nil {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"results": collections,
			"total":   count,
		})
	}
	return fiber.ErrNotFound
}

func CollectionById(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	collection := model.Collection{}

	result := db.Debug().Where("id = ?", id).First(&collection)

	if result.Error == nil {
		return c.JSON(map[string]interface{}{
			"status": "ok",
			"result": collection,
		})
	}
	return fiber.ErrNotFound
}

func CollectionByShortUrl(c *fiber.Ctx, db *gorm.DB) error {
	shortUrl := c.Query("short_url")

	collection := model.Collection{}

	result := db.Debug().Where("short_url = ?", shortUrl).First(&collection)

	if result.Error == nil {
		return c.JSON(map[string]interface{}{
			"status": "ok",
			"result": collection,
		})
	}
	return fiber.ErrNotFound
}

func CollectionByIdChain(c *fiber.Ctx, db *gorm.DB) error {
	idChain := c.Query("collection_id_chain")

	collection := model.Collection{}

	result := db.Debug().Where("collection_id_chain = ?", idChain).First(&collection)

	if result.Error == nil {
		return c.JSON(map[string]interface{}{
			"status": "ok",
			"result": collection,
		})
	}
	return fiber.ErrNotFound
}
