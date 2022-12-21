package handler

import (
	"reame-service/database/model"
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

type ProjectPayload struct {
	ProjectSource string             `json:"projectSource"`
	ProjectList   []model.Collection `json:"projectList"`
}

func (ph *ProjectHandler) Create(c *fiber.Ctx) error {
	return collection.Create(c, ph.DB)
}
