package handler

import (
	"fmt"
	"net/http"
	"reame-service/database/model"
	"reame-service/handler/upload"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type ExCollectionHandler struct {
	DB *gorm.DB
}

type PayloadData struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	ImageBanner      string `json:"imageBanner"`
	ImageFeature     string `json:"imageFeature"`
	ImageAvatar      string `json:"imageAvatar"`
	Description      string `json:"description"`
	Slug             string `json:"slug"`
	MinterAddress    string `json:"minterAddress"`
	TypeOnsale       bool   `json:"typeOnsale"`
	TypeCollection   bool   `json:"typeCollection"`
	Hot              bool   `json:"hot"`
	OwnerId          int    `json:"ownerId"`
	BranchId         int    `json:"branchId"`
	Instragram       string `json:"instragram"`
	Facebook         string `json:"facebook"`
	Active           string `json:"active"`
	Telegram         string `json:"telegram"`
	Discord          string `json:"discord"`
	Twitter          string `json:"twitter"`
	TokenType        string `json:"tokenType"`
	TermAndCondition string `json:"termAndCondition"`
	Status           string `json:"status"`
	ChainName        string `json:"chainName"`
}

func (ph *ExCollectionHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *ExCollectionHandler) PostCreateNewCollectionDetail(c *fiber.Ctx) error {
	userAddress := c.Locals("address").(string)
	payload := PayloadData{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var bannerLocation = ""
	if len(payload.ImageBanner) > 0 {
		bannerOutput, err := upload.AWSUpload(payload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-excollectionbanner", strings.ToLower(payload.ID)))
		if err == nil {
			bannerLocation = bannerOutput.Location
		}
	}

	var featureLocation = ""
	if len(payload.ImageFeature) > 0 {
		featureOutput, err := upload.AWSUpload(payload.ImageFeature, fmt.Sprintf("/%s/%s", "reame-excollectionfeature", strings.ToLower(payload.ID)))
		if err == nil {
			featureLocation = featureOutput.Location
		}
	}

	var avatarLocation = ""
	if len(payload.ImageAvatar) > 0 {
		avatarOutput, err := upload.AWSUpload(payload.ImageAvatar, fmt.Sprintf("/%s/%s", "reame-excollectionavatar", strings.ToLower(payload.ID)))
		if err == nil {
			avatarLocation = avatarOutput.Location
		}
	}

	collection := model.ExCollection{
		ID:               strings.ToLower(payload.ID),
		Title:            payload.Title,
		Description:      payload.Description,
		ImageBanner:      bannerLocation,
		ImageFeature:     featureLocation,
		ImageAvatar:      avatarLocation,
		TokenType:        payload.TokenType,
		Slug:             strings.ToLower(payload.Slug),
		Facebook:         payload.Facebook,
		Instragram:       payload.Instragram,
		Telegram:         payload.Telegram,
		Twitter:          payload.Twitter,
		Discord:          payload.Discord,
		Active:           payload.Active,
		TermAndCondition: payload.TermAndCondition,
		Status:           payload.Status,
		CreatedBy:        userAddress,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := ph.DB.Create(&collection).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *ExCollectionHandler) PutUpdateCollectionDetail(c *fiber.Ctx) error {
	userAddress := c.Locals("address").(string)
	payload := PayloadData{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var bannerLocation = ""
	if len(payload.ImageBanner) > 0 {
		bannerOutput, err := upload.AWSUpload(payload.ImageBanner, fmt.Sprintf("/%s/%s", "reame-excollectionbanner", strings.ToLower(payload.ID)))
		if err == nil {
			bannerLocation = bannerOutput.Location
		}
	}

	var featureLocation = ""
	if len(payload.ImageFeature) > 0 {
		featureOutput, err := upload.AWSUpload(payload.ImageFeature, fmt.Sprintf("/%s/%s", "reame-excollectionfeature", strings.ToLower(payload.ID)))
		if err == nil {
			featureLocation = featureOutput.Location
		}
	}

	var avatarLocation = ""
	if len(payload.ImageAvatar) > 0 {
		avatarOutput, err := upload.AWSUpload(payload.ImageAvatar, fmt.Sprintf("/%s/%s", "reame-excollectionavatar", strings.ToLower(payload.ID)))
		if err == nil {
			avatarLocation = avatarOutput.Location
		}
	}

	collection := model.ExCollection{
		Title:            payload.Title,
		Description:      payload.Description,
		ImageBanner:      bannerLocation,
		ImageFeature:     featureLocation,
		ImageAvatar:      avatarLocation,
		Slug:             strings.ToLower(payload.Slug),
		Facebook:         payload.Facebook,
		Instragram:       payload.Instragram,
		TokenType:        payload.TokenType,
		Telegram:         payload.Telegram,
		Twitter:          payload.Twitter,
		Active:           payload.Active,
		Discord:          payload.Discord,
		TermAndCondition: payload.TermAndCondition,
		Status:           payload.Status,
		UpdateBy:         userAddress,
		UpdatedAt:        time.Now(),
	}

	err := ph.DB.Model(&model.Collection{}).Where("ID = ?", strings.ToLower(payload.ID)).Updates(&collection).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *ExCollectionHandler) GetIsExist(c *fiber.Ctx) error {
	var collection model.Collection
	ref := c.Params("ref")

	query := strings.TrimSpace(strings.ToLower(ref))

	db := ph.DB.Model(&model.Collection{})

	var err error
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if re.MatchString(ref) {
		err = db.First(&collection, "id = ?", query).Error
	} else {
		err = db.First(&collection, "slug = ?", query).Error
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "ok",
			"result": false,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": true,
	})
}

func (ch *ExCollectionHandler) GetCollections(c *fiber.Ctx) error {
	pageQuery := c.Query("page", "1")
	pageSizeQuery := c.Query("page_size", "10")
	active := c.Query("active")

	collections := []model.Collection{}

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

	result := ch.DB.Offset(offset).Limit(pageSize)

	if len(active) > 0 {
		result.Where("active = ?", active)
	}

	result.Find(&collections)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": collections,
	})
}

func (ph *ExCollectionHandler) GetCollectionSingleDetail(c *fiber.Ctx) error {
	var collection model.Collection
	ref := c.Params("ref")

	collectionQuery := strings.TrimSpace(strings.ToLower(ref))

	db := ph.DB.Model(&model.Collection{})

	var err error
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if re.MatchString(collectionQuery) {
		err = db.Preload("Owner").First(&collection, "id = ?", collectionQuery).Error
	} else {
		err = db.Preload("Owner").First(&collection, "slug = ?", collectionQuery).Error
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"result": collection,
	})
}
