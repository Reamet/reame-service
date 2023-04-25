package collection

import (
	"fmt"
	"net/http"
	"reame-service/database/model"
	"reame-service/handler/upload"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreatePayload struct {
	CollectionIdChain      string `json:"collection_id_chain"`
	CollectionProfileImage string `json:"collection_profile_image"`
	CollectionCoverImage   string `json:"collection_cover_image"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	ShortUrl               string `json:"short_url"`
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

	uuidProfile := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var logoProfile = ""
	if len(bodyPayload.CollectionProfileImage) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.CollectionProfileImage, fmt.Sprintf("/%s/%s", "collection", uuidProfile))
		if err == nil {
			logoProfile = logoOutput.Location
		}
	}

	uuidWithHyphenCover := uuid.New()

	uuidCover := strings.Replace(uuidWithHyphenCover.String(), "-", "", -1)

	var logoCover = ""
	if len(bodyPayload.CollectionCoverImage) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.CollectionCoverImage, fmt.Sprintf("/%s/%s", "collection", uuidCover))
		if err == nil {
			logoCover = logoOutput.Location
		}
	}

	collectionDatabasePayload := model.Collection{
		CollectionIdChain:      bodyPayload.CollectionIdChain,
		CollectionProfileImage: logoProfile,
		CollectionCoverImage:   logoCover,
		Name:                   bodyPayload.Name,
		Description:            bodyPayload.Description,
		ShortUrl:               bodyPayload.ShortUrl,
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

type CreateTrendingPayload struct {
	Ids string `json:"ids"`
}

func CreateTrendingCollection(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := CreateTrendingPayload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	databasePayload := model.TrendingCollection{
		Ids:       bodyPayload.Ids,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := db.Create(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func UpdateTrendingCollection(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := CreateTrendingPayload{}
	currentTime := time.Now()

	id, errorId := strconv.Atoi(c.Params("id"))

	if errorId != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorId.Error(),
		})
	}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	databasePayload := model.TrendingCollection{
		Ids:       bodyPayload.Ids,
		UpdatedAt: currentTime,
	}

	if err := db.Where("id = ?", id).Updates(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func CreateFeaturedCollection(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := CreateTrendingPayload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	databasePayload := model.FeaturedCollection{
		Ids:       bodyPayload.Ids,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := db.Create(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func UpdateFeaturedCollection(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := CreateTrendingPayload{}
	currentTime := time.Now()

	id, errorId := strconv.Atoi(c.Params("id"))

	if errorId != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorId.Error(),
		})
	}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	databasePayload := model.FeaturedCollection{
		Ids:       bodyPayload.Ids,
		UpdatedAt: currentTime,
	}

	if err := db.Where("id = ?", id).Updates(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
