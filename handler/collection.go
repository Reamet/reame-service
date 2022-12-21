package handler

import (
	"reame-service/handler/collection"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProjectHandler struct {
	DB *gorm.DB
}

func (ph *ProjectHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *ProjectHandler) Create(c *fiber.Ctx) error {
	return collection.Create(c, ph.DB)
}
