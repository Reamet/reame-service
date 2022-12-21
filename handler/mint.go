package handler

import (
	"reame-service/handler/mint"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MintHandler struct {
	DB *gorm.DB
}

func (ph *MintHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *MintHandler) Create(c *fiber.Ctx) error {
	return mint.Create(c, ph.DB)
}

func (ph *MintHandler) Update(c *fiber.Ctx) error {
	return mint.Update(c, ph.DB)
}

func (ph *MintHandler) MintLists(c *fiber.Ctx) error {
	return mint.MintLists(c, ph.DB)
}

func (ph *MintHandler) GetMintById(c *fiber.Ctx) error {
	return mint.GetMintById(c, ph.DB)
}
