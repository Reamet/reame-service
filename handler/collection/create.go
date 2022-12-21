package collection

import (
	"fmt"
	"reame-service/database/model"
	"reame-service/handler/upload"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreatePayload struct {
	CollectionProfileImage string `json:"collection_profile_image"`
	CollectionCoverImage   string `json:"collection_cover_image"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	ShortUrl               string `json:"short_url"`
	Category               string `json:"category"`
	Address                string `json:"address"`
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

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var logoProfile = ""
	if len(bodyPayload.CollectionProfileImage) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.CollectionProfileImage, fmt.Sprintf("/%s/%s", "collection", uuid))
		if err == nil {
			logoProfile = logoOutput.Location
		}
	}

	uuid1 := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var logoCover = ""
	if len(bodyPayload.CollectionProfileImage) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.CollectionProfileImage, fmt.Sprintf("/%s/%s", "collection", uuid1))
		if err == nil {
			logoCover = logoOutput.Location
		}
	}

	collectionDatabasePayload := model.Collection{
		CollectionProfileImage: logoProfile,
		CollectionCoverImage:   logoCover,
		Name:                   bodyPayload.Name,
		Description:            bodyPayload.Description,
		ShortUrl:               bodyPayload.ShortUrl,
		Category:               bodyPayload.Category,
		Address:                bodyPayload.Address,
		Website:                bodyPayload.Website,
		Facebook:               bodyPayload.Facebook,
		Twitter:                bodyPayload.Twitter,
		Discord:                bodyPayload.Discord,
		Telegram:               bodyPayload.Telegram,
		Medium:                 bodyPayload.Medium,
		CreatedAt:              currentTime,
		UpdatedAt:              currentTime,
	}

	err := db.Debug().Create(&collectionDatabasePayload).Error

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
