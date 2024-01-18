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

type BranchHandler struct {
	DB *gorm.DB
}

type payloadBranch struct {
	OwnerId     int    `json:"ownerId"`
	Title       string `json:"title"`
	ImageBanner string `json:"imageBanner"`
	Description string `json:"description"`
}

func (ph *BranchHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *BranchHandler) CreateBranch(c *fiber.Ctx) error {
	payload := payloadBranch{}

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var BranchBannerLocation = ""
	if len(payload.ImageBanner) > 0 {
		BranchBannerOutput, err := upload.AWSUpload(payload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-branchbanner", strings.ToLower(uuid)))
		if err == nil {
			BranchBannerLocation = BranchBannerOutput.Location
		}
	}

	Branch := model.Branch{
		OwnerId:     payload.OwnerId,
		Title:       payload.Title,
		ImageBanner: BranchBannerLocation,
		Description: payload.Description,
		CreatedAt:   time.Now(),
	}

	var count int64

	titleExist := ph.DB.Model(&model.Branch{}).Where("title = ? AND owner_id = ?", payload.Title, payload.OwnerId).First(&Branch)
	titleExist.Offset(-1).Count(&count)

	if count > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "title name already exist.",
		})
	}

	err := ph.DB.Model(&model.Branch{}).Create(&Branch).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *BranchHandler) UpdateBranch(c *fiber.Ctx) error {
	payload := payloadBranch{}
	branchId := c.Params("branchId")

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	Branch := model.Branch{}

	var count int64

	findBranchID := ph.DB.Model(&model.Branch{}).Where("id = ?", branchId).First(&Branch)
	findBranchID.Offset(-1).Count(&count)

	if count <= 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "record not found.",
		})
	}

	titleExist := ph.DB.Model(&model.Branch{}).Where("title = ? AND owner_id = ?", payload.Title, payload.OwnerId).First(&Branch)
	titleExist.Offset(-1).Count(&count)

	if count > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "title name already exist.",
		})
	}

	var BranchBannerLocation = ""
	if len(payload.ImageBanner) > 0 {

		uuidWithHyphen := uuid.New()
		uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

		BranchBannerOutput, err := upload.AWSUpload(payload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-branchbanner", strings.ToLower(uuid)))
		if err == nil {
			BranchBannerLocation = BranchBannerOutput.Location
		}
	}

	BranchUpdate := model.Branch{
		OwnerId:     payload.OwnerId,
		Title:       payload.Title,
		ImageBanner: BranchBannerLocation,
		Description: payload.Description,
		UpdatedAt:   time.Now(),
	}

	err := ph.DB.Model(&model.Branch{}).Where("id = ?", branchId).Updates(&BranchUpdate).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ch *BranchHandler) GetBranchAll(c *fiber.Ctx) error {
	pageQuery := c.Query("page", "1")
	pageSizeQuery := c.Query("page_size", "10")

	branch := []model.Branch{}

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

	db := ch.DB.Model(&model.Branch{})

	var count int64

	result := db.Preload("Owner").Offset(offset).Limit(pageSize).Find(&branch)
	result.Offset(-1).Count(&count)

	err := result.Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": branch,
		"total":   count,
	})
}

func (ch *BranchHandler) GetBranchById(c *fiber.Ctx) error {
	branchId := c.Params("branchId")

	branch := model.Branch{}

	db := ch.DB.Model(&model.Branch{})

	result := db.Where("id = ?", branchId).Find(&branch)

	err := result.Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"result": branch,
	})
}

func (ch *BranchHandler) GetBranchByOwnerId(c *fiber.Ctx) error {
	ownerId := c.Params("ownerId")

	branch := []model.Branch{}

	db := ch.DB.Model(&model.Branch{})

	result := db.Where("owner_id = ?", ownerId).Find(&branch)

	err := result.Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": branch,
	})
}
