package collection

import (
	"fmt"
	"reame-service/database/model"
	"reame-service/handler/upload"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UpdatePayload struct {
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

	collectionModel := model.Collection{}

	db.Debug().Where("ID = ?", id).First(&collectionModel)

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var logoProfile = ""
	if len(collectionModel.CollectionProfileImage) > 0 {
		logoProfile = collectionModel.CollectionProfileImage
	}

	if len(bodyPayload.CollectionProfileImage) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.CollectionProfileImage, fmt.Sprintf("/%s/%s", "collection", uuid))
		if err == nil {
			logoProfile = logoOutput.Location
		}
	}

	var logoCover = ""
	if len(collectionModel.CollectionCoverImage) > 0 {
		logoCover = collectionModel.CollectionCoverImage
	}

	if len(bodyPayload.CollectionCoverImage) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.CollectionCoverImage, fmt.Sprintf("/%s/%s", "collection", uuid))
		if err == nil {
			logoCover = logoOutput.Location
		}
	}

	databasePayload := map[string]interface{}{
		"collection_profile_image": logoProfile,
		"collection_cover_image":   logoCover,
		"name":                     bodyPayload.Name,
		"description":              bodyPayload.Description,
		"short_url":                bodyPayload.ShortUrl,
		"category":                 bodyPayload.Category,
		"website":                  bodyPayload.Website,
		"facebook":                 bodyPayload.Facebook,
		"twitter":                  bodyPayload.Twitter,
		"discord":                  bodyPayload.Discord,
		"telegram":                 bodyPayload.Telegram,
		"medium":                   bodyPayload.Medium,
		"updated_at":               currentTime,
	}

	collectionResult := db.Debug().Where("ID = ?", id).First(&collectionModel).Updates(&databasePayload)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": fmt.Sprintf("Row Affected By : %s row", strconv.FormatInt(collectionResult.RowsAffected, 10)),
	})
}
