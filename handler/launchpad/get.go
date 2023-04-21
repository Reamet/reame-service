package launchpad

import (
	"net/http"
	"reame-service/database/model"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetLaunchPadAll(c *fiber.Ctx, db *gorm.DB) error {
	pageQuery := c.Query("page", "1")
	pageSizeQuery := c.Query("page_size", "10")
	hot := c.Query("hot")
	status := c.Query("status")
	saleStatus := c.Query("salestatus")
	saleType := c.Query("saletype")

	launchpads := []model.Launchpad{}

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

	var count int64

	result := db.Debug().Offset(offset).Limit(pageSize)

	if len(hot) > 0 {
		result.Where("hot = ?", hot)
	}

	if len(status) > 0 {
		result.Where("status = ?", status)
	}

	if len(saleStatus) > 0 {
		result.Where("sale_status = ?", saleStatus)
	}

	if len(saleType) > 0 {
		result.Where("sale_type = ?", saleType)
	}

	result.Find(&launchpads)
	result.Offset(-1).Count(&count)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"results": launchpads,
		"total":   count,
	})
}

func GetLaunchPadById(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	launchpad := model.Launchpad{}

	launchpadQuery := strings.TrimSpace(strings.ToLower(id))

	result := db.Debug().Where("id = ?", launchpadQuery).First(&launchpad)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": launchpad,
	})
}

func GetLaunchPadBySlug(c *fiber.Ctx, db *gorm.DB) error {
	slug := c.Params("slug")
	launchpad := model.Launchpad{}

	launchpadQuery := strings.TrimSpace(slug)

	result := db.Debug().Where("slug = ?", launchpadQuery).First(&launchpad)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"result": launchpad,
	})
}
