package mint

import (
	"reame-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreatePayload struct {
	MintImage   string `json:"mint_image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Collection  string `json:"collection"`
	Address     string `json:"address"`
	Royalties   string `json:"royalties"`
}

func Create(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := CreatePayload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	mintDatabasePayload := model.Mint{
		MintImage:   bodyPayload.MintImage,
		Name:        bodyPayload.Name,
		Description: bodyPayload.Description,
		Collection:  bodyPayload.Collection,
		Address:     bodyPayload.Address,
		Royalties:   bodyPayload.Royalties,
		CreatedAt:   currentTime,
	}

	err := db.Debug().Create(&mintDatabasePayload).Error

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
