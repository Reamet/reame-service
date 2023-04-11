package handler

import (
	"reame-service/handler/launchpad"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LaunchpadHandler struct {
	DB *gorm.DB
}

func (ph *LaunchpadHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *LaunchpadHandler) Create(c *fiber.Ctx) error {
	return launchpad.Create(c, ph.DB)
}

func (ph *LaunchpadHandler) Update(c *fiber.Ctx) error {
	return launchpad.Update(c, ph.DB)
}

func (ph *LaunchpadHandler) GetLaunchPadAll(c *fiber.Ctx) error {
	return launchpad.GetLaunchPadAll(c, ph.DB)
}

func (ph *LaunchpadHandler) GetLaunchPadById(c *fiber.Ctx) error {
	return launchpad.GetLaunchPadById(c, ph.DB)
}

func (ph *LaunchpadHandler) GetLaunchPadBySlug(c *fiber.Ctx) error {
	return launchpad.GetLaunchPadBySlug(c, ph.DB)
}
