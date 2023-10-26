package featuredlaunchpad

import (
	"net/http"
	"reame-service/database/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FeaturedLaunchpadHandler struct {
	DB *gorm.DB
}

func (hd *FeaturedLaunchpadHandler) GetAllFeaturedLaunchpad(c *fiber.Ctx) error {
	launchpad := []model.FeaturedLaunchpad{}

	result := hd.DB.Order("id DESC").Find(&launchpad)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": launchpad,
	})
}

func (ph *FeaturedLaunchpadHandler) CreateFeaturedLaunchpad(c *fiber.Ctx) error {
	type Payload struct {
		Ids string `json:"ids"`
	}
	bodyPayload := Payload{}
	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	databasePayload := model.FeaturedLaunchpad{
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

func (ph *FeaturedLaunchpadHandler) UpdateFeaturedLaunchpad(c *fiber.Ctx) error {
	type Payload struct {
		Ids string `json:"ids"`
	}
	bodyPayload := Payload{}
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

	databasePayload := model.FeaturedLaunchpad{
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
