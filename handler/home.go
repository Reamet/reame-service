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

type HomeHandler struct {
	DB *gorm.DB
}

func (ph *HomeHandler) Init(db *gorm.DB) {
	ph.DB = db
}

type CreateHomePayload struct {
	Title       string `json:"title"`
	Subtitle    string `json:"Subtitle"`
	ImageBanner string `json:"imageBanner"`
	ButtonTitle string `json:"buttonTitle"`
	NftIds      string `json:"nftIds"`
}

func (ph *HomeHandler) Create(c *fiber.Ctx) error {
	bodyPayload := CreateHomePayload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	uuidWithHyphen := uuid.New()

	uuidBanner := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var imageBanner = ""
	if len(bodyPayload.ImageBanner) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-banner", uuidBanner))
		if err == nil {
			imageBanner = logoOutput.Location
		}
	}

	databasePayload := model.Home{
		Title:       bodyPayload.Title,
		Subtitle:    bodyPayload.Subtitle,
		ImageBanner: imageBanner,
		ButtonTitle: bodyPayload.ButtonTitle,
		NftIds:      bodyPayload.NftIds,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	if err := ph.DB.Create(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *HomeHandler) Update(c *fiber.Ctx) error {
	bodyPayload := CreateHomePayload{}
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

	uuidWithHyphen := uuid.New()

	uuidBanner := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	var imageBanner = ""
	if len(bodyPayload.ImageBanner) > 0 {
		logoOutput, err := upload.AWSUpload(bodyPayload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-banner", uuidBanner))
		if err == nil {
			imageBanner = logoOutput.Location
		}
	}

	databasePayload := model.Home{
		Title:       bodyPayload.Title,
		Subtitle:    bodyPayload.Subtitle,
		ImageBanner: imageBanner,
		ButtonTitle: bodyPayload.ButtonTitle,
		NftIds:      bodyPayload.NftIds,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	if err := ph.DB.Where("id = ?", id).Updates(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *HomeHandler) GetHomeById(c *fiber.Ctx) error {
	id, errorId := strconv.Atoi(c.Params("id"))
	if errorId != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorId.Error(),
		})
	}

	home := model.Home{}

	result := ph.DB.Where("id = ?", id).First(&home)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": home,
	})
}

func (ph *HomeHandler) GetAllHome(c *fiber.Ctx) error {
	pageQuery := c.Query("page", "1")
	pageSizeQuery := c.Query("page_size", "10")

	home := []model.Home{}

	page, _ := strconv.Atoi(pageQuery)
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(pageSizeQuery)
	var count int64

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	result := ph.DB.Limit(pageSize).Offset(offset).Find(&home)
	result.Offset(-1).Count(&count)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"results": home,
		"total":   count,
	})
}
