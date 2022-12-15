package project

import (
	"bsc-scan-data-service/database/model"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProjectListPayload struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Information   string   `json:"information"`
	Logo          string   `json:"logo"`
	Website       string   `json:"website"`
	ProjectSource string   `json:"projectSource"`
	Telegram      *string  `json:"telegram"`
	Twitter       *string  `json:"twitter"`
	Discord       *string  `json:"discord"`
	Email         *string  `json:"email"`
	Facebook      *string  `json:"facebook"`
	Instagram     *string  `json:"instagram"`
	ReferredId    *float64 `json:"id"`
}

type ProjectPayload struct {
	ProjectSource string               `json:"projectSource"`
	ProjectList   []ProjectListPayload `json:"projectList"`
}

func ProjectCreateJson(c *fiber.Ctx, db *gorm.DB) error {

	bodyPayload := ProjectPayload{}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	projectResponse := []model.Project{}

	for _, project := range bodyPayload.ProjectList {

		projectResult := model.Project{}

		result := db.Debug().Where("referred_id = ?", project.ReferredId).First(&projectResult)

		if result.RowsAffected == 0 {

			databasePayload := model.Project{
				ReferredId:  project.ReferredId,
				Title:       project.Title,
				Description: project.Description,
				Source:      bodyPayload.ProjectSource,
				Logo:        project.Logo,
				Website:     project.Website,
				Information: project.Information,
				Telegram:    project.Telegram,
				Twitter:     project.Twitter,
				Discord:     project.Discord,
				Email:       project.Email,
				Facebook:    project.Facebook,
				Instagram:   project.Instagram,
			}

			err := db.Debug().Create(&databasePayload).Error
			projectResponse = append(projectResponse, databasePayload)
			if err != nil {
				return err
			}

		} else {

			databasePayload := model.Project{
				ReferredId:  project.ReferredId,
				Title:       project.Title,
				Description: project.Description,
				Source:      bodyPayload.ProjectSource,
				Logo:        project.Logo,
				Website:     project.Website,
				Information: project.Information,
				Telegram:    project.Telegram,
				Twitter:     project.Twitter,
				Discord:     project.Discord,
				Email:       project.Email,
				Facebook:    project.Facebook,
				Instagram:   project.Instagram,
			}

			err := result.Debug().Updates(&databasePayload).Error
			projectResponse = append(projectResponse, databasePayload)
			if err != nil {
				return err
			}
		}
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"results": projectResponse,
	})
}

func ProjectCreate(c *fiber.Ctx, db *gorm.DB) error {

	bodyPayload := ProjectListPayload{}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	databasePayload := model.Project{
		ReferredId:  bodyPayload.ReferredId,
		Title:       bodyPayload.Title,
		Description: bodyPayload.Description,
		Source:      bodyPayload.ProjectSource,
		Logo:        bodyPayload.Logo,
		Website:     bodyPayload.Website,
		Information: bodyPayload.Information,
		Telegram:    bodyPayload.Telegram,
		Twitter:     bodyPayload.Twitter,
		Discord:     bodyPayload.Discord,
		Email:       bodyPayload.Email,
		Facebook:    bodyPayload.Facebook,
		Instagram:   bodyPayload.Instagram,
	}

	err := db.Debug().Create(&databasePayload).Error
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"id":     databasePayload.ID,
	})
}

func ProjectUpdate(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	bodyPayload := ProjectListPayload{}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	databasePayload := model.Project{
		Title:       bodyPayload.Title,
		Description: bodyPayload.Description,
		Source:      bodyPayload.ProjectSource,
		Logo:        bodyPayload.Logo,
		Website:     bodyPayload.Website,
		Information: bodyPayload.Information,
		Telegram:    bodyPayload.Telegram,
		Twitter:     bodyPayload.Twitter,
		Discord:     bodyPayload.Discord,
		Email:       bodyPayload.Email,
		Facebook:    bodyPayload.Facebook,
		Instagram:   bodyPayload.Instagram,
	}

	projectModel := model.Project{}

	poolResult := db.Debug().Where("ID = ?", id).First(&projectModel).Updates(&databasePayload)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": fmt.Sprintf("Row Affected By : %s row", strconv.FormatInt(poolResult.RowsAffected, 10)),
	})
}
