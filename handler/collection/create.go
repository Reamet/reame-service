package proposal

import (
	"reame-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreatePayload struct {
	CollectionProfileImage string `json:"collection_profile_image"`
	CollectionCoverImage   string `json:"collection_cover_image"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	ShortUrl               string `json:"short_url"`
	Category               string `json:"category"`
	Website                string `json:"website"`
	Facebook               string `json:"facebook"`
	Twitter                string `json:"twitter"`
	Discord                string `json:"discord"`
	Telegram               string `json:"telegram"`
	Medium                 string `json:"medium"`
}

func Create(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := CreatePayload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	proposalDatabasePayload := model.Collection{
		CollectionProfileImage: bodyPayload.CollectionProfileImage,
		CollectionCoverImage:   bodyPayload.CollectionCoverImage,
		Name:                   bodyPayload.Name,
		Description:            bodyPayload.Description,
		ShortUrl:               bodyPayload.ShortUrl,
		Category:               bodyPayload.Category,
		Website:                bodyPayload.Website,
		Facebook:               bodyPayload.Facebook,
		Twitter:                bodyPayload.Twitter,
		Discord:                bodyPayload.Discord,
		Telegram:               bodyPayload.Telegram,
		Medium:                 bodyPayload.Medium,
		CreatedAt:              currentTime,
	}

	err := db.Debug().Create(&proposalDatabasePayload).Error

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
