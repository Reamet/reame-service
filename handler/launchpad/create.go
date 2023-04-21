package launchpad

import (
	"fmt"
	"net/http"
	"reame-service/database/model"
	"reame-service/handler/upload"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreatePayload struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	ImageBanner      string    `json:"imageBanner"`
	ImageFeature     string    `json:"imageFeature"`
	ImageAvatar      string    `json:"imageAvatar"`
	ChainName        string    `json:"chainName"`
	LaunchpadAddress string    `json:"launchpadAddress"`
	Hot              *bool     `json:"hot"`
	Slug             string    `json:"slug"`
	Instragram       string    `json:"instragram"`
	Twitter          string    `json:"twitter"`
	Facebook         string    `json:"facebook"`
	Telegram         string    `json:"telegram"`
	Discord          string    `json:"discord"`
	CreatedBy        string    `json:"createdBy"`
	UpdateBy         string    `json:"updateBy"`
	Active           *bool     `json:"active"`
	TermAndCondition string    `json:"termAndCondition"`
	Status           string    `json:"status"`
	SaleStatus       string    `json:"saleStatus"`
	SaleType         string    `json:"saleType"`
	StartDate        time.Time `json:"startDate,omitempty"`
	EndDate          time.Time `json:"endDate,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
}

func Create(c *fiber.Ctx, db *gorm.DB) error {
	payload := CreatePayload{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var bannerLocation = ""
	if len(payload.ImageBanner) > 0 {
		bannerOutput, err := upload.AWSUpload(payload.ImageBanner, fmt.Sprintf("/%s/%s", "reamelaunchpadbanner", strings.ToLower(payload.ID)))
		if err == nil {
			bannerLocation = bannerOutput.Location
		}
	}

	var featureLocation = ""
	if len(payload.ImageFeature) > 0 {
		featureOutput, err := upload.AWSUpload(payload.ImageFeature, fmt.Sprintf("/%s/%s", "reamelaunchpadfeature", strings.ToLower(payload.ID)))
		if err == nil {
			featureLocation = featureOutput.Location
		}
	}

	var avatarLocation = ""
	if len(payload.ImageAvatar) > 0 {
		avatarOutput, err := upload.AWSUpload(payload.ImageAvatar, fmt.Sprintf("/%s/%s", "reamelaunchpadavatar", strings.ToLower(payload.ID)))
		if err == nil {
			avatarLocation = avatarOutput.Location
		}
	}

	launchpad := model.Launchpad{
		ID:               strings.ToLower(payload.ID),
		Title:            payload.Title,
		Description:      payload.Description,
		ImageBanner:      bannerLocation,
		ImageFeature:     featureLocation,
		ImageAvatar:      avatarLocation,
		ChainName:        payload.ChainName,
		LaunchpadAddress: payload.LaunchpadAddress,
		Hot:              payload.Hot,
		Slug:             strings.ToLower(payload.Slug),
		Facebook:         payload.Facebook,
		Instragram:       payload.Instragram,
		Telegram:         payload.Telegram,
		Twitter:          payload.Twitter,
		Discord:          payload.Discord,
		Active:           payload.Active,
		TermAndCondition: payload.TermAndCondition,
		Status:           payload.Status,
		SaleStatus:       payload.SaleStatus,
		SaleType:         payload.SaleType,
		StartDate:        payload.StartDate,
		EndDate:          payload.EndDate,
		CreatedBy:        payload.CreatedBy,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := db.Create(&launchpad).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}
