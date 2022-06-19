package handler

import (
	project_pool "bsc-scan-data-service/handler/project-pool"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProjectPoolHandler struct {
	DB *gorm.DB
}

func (pph *ProjectPoolHandler) Init(db *gorm.DB) {
	pph.DB = db
}

func (pph *ProjectPoolHandler) ProjectPoolCreate(c *fiber.Ctx) error {
	return project_pool.ProjectPoolCreate(c, pph.DB)
}

func (pph *ProjectPoolHandler) ProjectPoolList(c *fiber.Ctx) error {
	return project_pool.ProjectPoolList(c, pph.DB)
}

func (pph *ProjectPoolHandler) ProjectPoolUpdate(c *fiber.Ctx) error {
	return project_pool.ProjectPoolUpdate(c, pph.DB)
}

func (pph *ProjectPoolHandler) ProjectPoolById(c *fiber.Ctx) error {
	return project_pool.ProjectPoolById(c, pph.DB)
}
