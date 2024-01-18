package handler

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

type OwnerHandler struct {
	DB *gorm.DB
}

type payloadOwner struct {
	Title       string `json:"title"`
	ImageBanner string `json:"imageBanner"`
	Description string `json:"description"`
}

func (ph *OwnerHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *OwnerHandler) CreateOwner(c *fiber.Ctx) error {
	payload := payloadOwner{}

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var ownerBannerLocation = ""
	if len(payload.ImageBanner) > 0 {
		OwnerBannerOutput, err := upload.AWSUpload(payload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-ownerbanner", strings.ToLower(uuid)))
		if err == nil {
			ownerBannerLocation = OwnerBannerOutput.Location
		}
	}

	owner := model.Owner{
		Title:       payload.Title,
		ImageBanner: ownerBannerLocation,
		Description: payload.Description,
		CreatedAt:   time.Now(),
	}

	var count int64

	titleExist := ph.DB.Model(&model.Owner{}).Where("title = ?", payload.Title).First(&owner)
	titleExist.Offset(-1).Count(&count)

	if count > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "title name already exist.",
		})
	}

	err := ph.DB.Model(&model.Owner{}).Create(&owner).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *OwnerHandler) UpdateOwner(c *fiber.Ctx) error {
	payload := payloadOwner{}
	ownerId := c.Params("ownerId")

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	Owner := model.Owner{}

	var count int64

	findOwnerID := ph.DB.Model(&model.Owner{}).Where("id = ?", ownerId).First(&Owner)
	findOwnerID.Offset(-1).Count(&count)

	if count <= 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "record not found.",
		})
	}

	titleExist := ph.DB.Model(&model.Owner{}).Where("title = ? ", payload.Title).First(&Owner)
	titleExist.Offset(-1).Count(&count)

	if count > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "title name already exist.",
		})
	}

	var OwnerBannerLocation = ""
	if len(payload.ImageBanner) > 0 {

		uuidWithHyphen := uuid.New()
		uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

		OwnerBannerOutput, err := upload.AWSUpload(payload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-ownerbanner", strings.ToLower(uuid)))
		if err == nil {
			OwnerBannerLocation = OwnerBannerOutput.Location
		}
	}

	OwnerUpdate := model.Owner{
		Title:       payload.Title,
		ImageBanner: OwnerBannerLocation,
		Description: payload.Description,
		UpdatedAt:   time.Now(),
	}

	err := ph.DB.Model(&model.Owner{}).Where("id = ?", ownerId).Updates(&OwnerUpdate).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ch *OwnerHandler) GetOwnerAll(c *fiber.Ctx) error {
	pageQuery := c.Query("page", "1")
	pageSizeQuery := c.Query("page_size", "10")

	owner := []model.Owner{}

	page, _ := strconv.Atoi(pageQuery)
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(pageSizeQuery)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	db := ch.DB.Model(&model.Owner{})

	var count int64

	result := db.Offset(offset).Limit(pageSize).Find(&owner)
	result.Offset(-1).Count(&count)

	err := result.Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": owner,
		"total":   count,
	})
}

func (ch *OwnerHandler) GetOwnerById(c *fiber.Ctx) error {
	ownerId := c.Params("ownerId")

	owner := model.Owner{}

	db := ch.DB.Model(&model.Owner{})

	result := db.Where("id = ?", ownerId).Find(&owner)

	err := result.Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"result": owner,
	})
}
