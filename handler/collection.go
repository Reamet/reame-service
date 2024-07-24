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

type CollectionHandler struct {
	DB *gorm.DB
}

type PayloadData struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	CollectionIdChain string `json:"collection_id_chain"`
	ImageBanner       string `json:"image_banner"`
	ImageFeature      string `json:"image_feature"`
	ImageAvatar       string `json:"image_avatar"`
	Description       string `json:"description"`
	Slug              string `json:"slug"`
	Hot               bool   `json:"hot"`
	OwnerId           int    `json:"owner_id"`
	BranchId          int    `json:"branch_id"`
	Instragram        string `json:"instragram"`
	Medium			  string `json:"medium"`
	Facebook          string `json:"facebook"`
	Active            string `json:"active"`
	Telegram          string `json:"telegram"`
	Discord           string `json:"discord"`
	Website           string `json:"website"`
	Twitter           string `json:"twitter"`
	TokenType         string `json:"token_type"`
	TermAndCondition  string `json:"term_and_condition"`
	Status            string `json:"status"`
	CreatedType       string `json:"created_type"`
	CreatedBy         string `json:"created_by"`
}

func (ph *CollectionHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *CollectionHandler) PostCreateNewCollectionDetail(c *fiber.Ctx) error {
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

	collection := model.Collection{
		ID:                strings.ToLower(payload.ID),
		CollectionIdChain: &payload.CollectionIdChain,
		Title:             payload.Title,
		Description:       payload.Description,
		ImageBanner:       bannerLocation,
		ImageFeature:      &featureLocation,
		ImageAvatar:       avatarLocation,
		OwnerId:           &payload.OwnerId,
		BranchId:          &payload.BranchId,
		TokenType:         &payload.TokenType,
		Slug:              strings.ToLower(payload.Slug),
		Facebook:          payload.Facebook,
		Instragram:        payload.Instragram,
		Telegram:          payload.Telegram,
		Twitter:           payload.Twitter,
		Discord:           payload.Discord,
		Website:           payload.Website,
		Active:            payload.Active,
		TermAndCondition:  payload.TermAndCondition,
		Status:            payload.Status,
		CreatedType:       payload.CreatedType,
		CreatedBy:         payload.CreatedBy,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
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

func (ph *CollectionHandler) PutUpdateCollectionDetail(c *fiber.Ctx) error {
	payload := PayloadData{}
	id := c.Params("id")

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

	collection := model.Collection{
		Title:             payload.Title,
		Description:       payload.Description,
		CollectionIdChain: &payload.CollectionIdChain,
		ImageBanner:       bannerLocation,
		ImageFeature:      &featureLocation,
		ImageAvatar:       avatarLocation,
		OwnerId:           &payload.OwnerId,
		BranchId:          &payload.BranchId,
		TokenType:         &payload.TokenType,
		Slug:              strings.ToLower(payload.Slug),
		Facebook:          payload.Facebook,
		Instragram:        payload.Instragram,
		Telegram:          payload.Telegram,
		Twitter:           payload.Twitter,
		Discord:           payload.Discord,
		Website:           payload.Website,
		Active:            payload.Active,
		TermAndCondition:  payload.TermAndCondition,
		Status:            payload.Status,
		UpdatedAt:         time.Now(),
	}

	err := ph.DB.Model(&model.Collection{}).Where("ID = ?", strings.ToLower(id)).Updates(&collection).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *CollectionHandler) GetIsExist(c *fiber.Ctx) error {
	collection := model.Collection{}
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

func (ch *CollectionHandler) GetCollections(c *fiber.Ctx) error {
	pageQuery := c.Query("page", "1")
	pageSizeQuery := c.Query("page_size", "10")
	active := c.Query("active")
	createby := c.Query("createby")

	collections := []model.Collection{}
	var count int64

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
	if len(createby) > 0 {
		result.Where("created_by = ?", createby)
	}

	result.Find(&collections)
	result.Offset(-1).Count(&count)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": collections,
		"total":   count,
	})
}

func (ph *CollectionHandler) GetCollectionSingleDetail(c *fiber.Ctx) error {
	collection := model.Collection{}
	ref := c.Params("ref")

	collectionQuery := strings.TrimSpace(strings.ToLower(ref))

	db := ph.DB.Model(&model.Collection{})

	var err error
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if re.MatchString(collectionQuery) {
		err = db.First(&collection, "id = ?", collectionQuery).Error
	} else {
		err = db.First(&collection, "slug = ?", collectionQuery).Error
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": collection,
	})
}

func (ph *CollectionHandler) CollectionByIdChain(c *fiber.Ctx) error {
	idChain := c.Query("collection_id_chain")

	collection := model.Collection{}

	result := ph.DB.Debug().Where("collection_id_chain = ?", idChain).First(&collection)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status": "ok",
		"result": collection,
	})
}

type CreateTrendingPayload struct {
	Ids string `json:"ids"`
}

func (ph *CollectionHandler) CreateTrendingCollection(c *fiber.Ctx) error {
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

	if err := ph.DB.Create(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *CollectionHandler) UpdateTrendingCollection(c *fiber.Ctx) error {
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

	if err := ph.DB.Where("id = ?", id).Updates(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *CollectionHandler) CreateFeaturedCollection(c *fiber.Ctx) error {
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

	if err := ph.DB.Create(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *CollectionHandler) UpdateFeaturedCollection(c *fiber.Ctx) error {
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

	if err := ph.DB.Where("id = ?", id).Updates(&databasePayload).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ph *CollectionHandler) GetAllTrendingCollection(c *fiber.Ctx) error {
	collection := []model.TrendingCollection{}

	result := ph.DB.Debug().Order("id DESC").Find(&collection)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": collection,
	})
}

func (ph *CollectionHandler) GetAllFeaturedCollection(c *fiber.Ctx) error {
	collection := []model.FeaturedCollection{}

	result := ph.DB.Debug().Order("id DESC").Find(&collection)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": collection,
	})
}
