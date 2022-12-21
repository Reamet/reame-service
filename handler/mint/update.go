package mint

import (
	"fmt"
	"reame-service/database/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdatePayload struct {
	MintImage   string `json:"mint_image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Collection  string `json:"collection"`
	Address     string `json:"address"`
	Royalties   string `json:"royalties"`
}

func Update(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	currentTime := time.Now()

	bodyPayload := UpdatePayload{}
	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	mintModel := model.Mint{}

	db.Debug().Where("ID = ?", id).First(&mintModel)

	databasePayload := map[string]interface{}{
		"mint_image":  bodyPayload.MintImage,
		"name":        bodyPayload.Name,
		"description": bodyPayload.Description,
		"collection":  bodyPayload.Collection,
		"address":     bodyPayload.Address,
		"royalties":   bodyPayload.Royalties,
		"updated_at":  currentTime,
	}

	collectionResult := db.Debug().Where("ID = ?", id).First(&mintModel).Updates(&databasePayload)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": fmt.Sprintf("Row Affected By : %s row", strconv.FormatInt(collectionResult.RowsAffected, 10)),
	})
}
