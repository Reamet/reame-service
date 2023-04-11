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

type UpdatePayload struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	ImageBanner      string    `json:"imageBanner"`
	ImageFeature     string    `json:"imageFeature"`
	ImageAvatar      string    `json:"imageAvatar"`
	ChainName        string    `json:"chainName"`
	LaunchpadAddress string    `json:"launchpadAddress"`
	Hot              bool      `json:"hot"`
	Slug             string    `json:"slug"`
	Instragram       string    `json:"instragram"`
	Twitter          string    `json:"twitter"`
	Facebook         string    `json:"facebook"`
	Telegram         string    `json:"telegram"`
	Discord          string    `json:"discord"`
	CreatedBy        string    `json:"createdBy"`
	UpdateBy         string    `json:"updateBy"`
	Active           string    `json:"active"`
	TermAndCondition string    `json:"termAndCondition"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
}

func Update(c *fiber.Ctx, db *gorm.DB) error {
	// userAddress := c.Locals("address").(string)
	payload := UpdatePayload{}

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
		Title:            payload.Title,
		Description:      payload.Description,
		ImageBanner:      bannerLocation,
		ImageFeature:     featureLocation,
		ImageAvatar:      avatarLocation,
		Slug:             strings.ToLower(payload.Slug),
		LaunchpadAddress: payload.LaunchpadAddress,
		Hot:              payload.Hot,
		Facebook:         payload.Facebook,
		Instragram:       payload.Instragram,
		ChainName:        payload.ChainName,
		Telegram:         payload.Telegram,
		Twitter:          payload.Twitter,
		Active:           payload.Active,
		Discord:          payload.Discord,
		TermAndCondition: payload.TermAndCondition,
		Status:           payload.Status,
		UpdateBy:         payload.CreatedBy,
		UpdatedAt:        time.Now(),
	}

	err := db.Debug().Model(&model.Launchpad{}).Where("ID = ?", strings.ToLower(payload.ID)).Updates(&launchpad).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "ok",
	})
}
