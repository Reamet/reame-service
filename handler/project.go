package handler

import (
	"bsc-scan-data-service/database/model"
	"bsc-scan-data-service/handler/project"

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
	ProjectSource string          `json:"projectSource"`
	ProjectList   []model.Project `json:"projectList"`
}

func (ph *ProjectHandler) ProjectCreate(c *fiber.Ctx) error {
	return project.ProjectCreate(c, ph.DB)
}

func (ph *ProjectHandler) ProjectLists(c *fiber.Ctx) error {
	return project.GetProjectLists(c, ph.DB)
}

func (ph *ProjectHandler) ProjectById(c *fiber.Ctx) error {
	return project.ProjectById(c, ph.DB)
}
