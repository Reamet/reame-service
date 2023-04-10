package launchpad

import (
	"fmt"
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

	result := db.Debug().Offset(offset).Limit(pageSize)
	if len(hot) > 0 {
		result.Where("hot = ?", hot)
	}

	if len(status) > 0 {
		result.Where("status = ?", status)
	}

	result.Find(&launchpads)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": launchpads,
	})
}

func GetLaunchPadById(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	launchpads := model.Launchpad{}

	launchpadQuery := strings.TrimSpace(strings.ToLower(id))
	fmt.Println(launchpadQuery)

	result := db.Debug().Where("id = ?", launchpadQuery).First(&launchpads)
	fmt.Println(result)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": launchpads,
	})
}

func GetLaunchPadBySlug(c *fiber.Ctx, db *gorm.DB) error {
	slug := c.Params("slug")
	launchpads := []model.Launchpad{}

	launchpadQuery := strings.TrimSpace(slug)

	result := db.Debug().Where("slug = ?", launchpadQuery).First(&launchpads)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": launchpads,
	})
}
