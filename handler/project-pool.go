package handler

import (
	"bsc-scan-data-service/database/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
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
	ProjectList					[]int64					`json:"projectList"`
}


func (pph *ProjectPoolHandler) ProjectPoolCreate(c *fiber.Ctx) error {
	bodyPayload := ProjectPoolCreatePayload{}

	currentTime := time.Now()

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
		UpdatedAt: currentTime,
		CreatedAt: currentTime,
	}

	err := pph.DB.Debug().Create(&databasePayload).Error

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (pph *ProjectPoolHandler) ProjectPoolList(c *fiber.Ctx) error {
	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		return err
	}

	projectPools := []model.ProjectPool{}

	result := pph.DB.Debug().Limit(10).Order("created_at desc").Offset(offset).Find(&projectPools)

	var count int64

	result.Debug().Offset(-1).Count(&count)

	return c.JSON(fiber.Map{
		"status": "ok",
		"project_pool_list": projectPools,
		"amount": count,
	})
}

type ProjectPoolUpdatePayload struct {
	ID									int							`json:"id"`
	Title								string					`json:"title"`
	SubTitle						string					`json:"subTitle"`
	Description					string					`json:"description"`
	Source							string					`json:"source"`
	StartDate						time.Time				`json:"startDate"`
	EndDate							time.Time				`json:"endDate"`
	ProjectList					pq.Int64Array		`json:"projectList"`
}

func (pph *ProjectPoolHandler) ProjectPoolUpdate(c *fiber.Ctx) error {
	bodyPayload := ProjectPoolUpdatePayload{}

	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	databasePayload := map[string]interface{} {
		"title": bodyPayload.Title,
		"sub_title": bodyPayload.SubTitle,
		"description": bodyPayload.Description,
		"source": bodyPayload.Source,
		"start_date": bodyPayload.StartDate,
		"end_date": bodyPayload.EndDate,
		"project_list": bodyPayload.ProjectList,
		"updated_at": currentTime,
	}

	pool := model.ProjectPool{}

	result := pph.DB.Debug().Where("ID = ?", bodyPayload.ID).First(&pool).Updates(&databasePayload)


		
	return c.JSON(fiber.Map{
		"status": "ok",
		"message": fmt.Sprintf("Row Affected By : %s row", strconv.FormatInt(result.RowsAffected, 10)),
	})
}