package handler

import (
	"bsc-scan-data-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProjectPoolHandler struct {
	DB *gorm.DB
}

func (pph *ProjectPoolHandler) Init(db *gorm.DB) {
	pph.DB = db
}

type ProjectPoolCreatePayload struct {
	Title								string					`json:"title"`
	SubTitle						string					`json:"subTitle"`
	Description					string					`json:"description"`
	Source							string					`json:"source"`
	StartDate						time.Time				`json:"startDate"`
	EndDate							time.Time				`json:"endDate"`
	ProjectList					[]uint					`json:"projectList"`
}


func (pph *ProjectPoolHandler) ProjectPoolCreate(c *fiber.Ctx) error {
	bodyPayload := ProjectPoolCreatePayload{}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	databasePayload := model.ProjectPool {
		Title: bodyPayload.Title,
		SubTitle: bodyPayload.SubTitle,
		Description: bodyPayload.Description,
		Source: bodyPayload.Source,
		StartDate: bodyPayload.StartDate,
		EndDate: bodyPayload.EndDate,
		ProjectList: bodyPayload.ProjectList,
	}

	err := pph.DB.Debug().Create(&databasePayload).Error

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}