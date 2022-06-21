package project_pool

import (
	"bsc-scan-data-service/database/model"
	project_pool_response_formatter "bsc-scan-data-service/handler/project-pool/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProjectPoolList(c *fiber.Ctx, db *gorm.DB) error {
	offset, err := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))
	status := c.Query("status")

	if err != nil {
		return err
	}

	projectPoolsResponse := []model.ProjectPoolResponse{}

	projectPools := []model.ProjectPool{}
	var count int64

	if status != "" {
		poolsResult := db.Debug().Where("status = ?", status).Preload("TierList", "deleted_at IS NULL").Limit(limit).Order("created_at desc").Offset(offset).Find(&projectPools)

		poolsResult.Debug().Offset(-1).Count(&count)

		for _, projectPool := range projectPools {

			project := []model.Project{}

			projectPoolList := []int64(projectPool.ProjectList)

			db.Debug().Where("id IN ?", projectPoolList).Find(&project)

			projectPoolsResponse = append(projectPoolsResponse, project_pool_response_formatter.ProjectPoolResponseFormatter(projectPool, project))
		}

		return c.JSON(fiber.Map{
			"status":            "ok",
			"project_pool_list": projectPoolsResponse,
			"amount":            count,
		})
	}

	poolsResult := db.Debug().Preload("TierList", "deleted_at IS NULL").Limit(limit).Order("created_at desc").Offset(offset).Find(&projectPools)

	poolsResult.Debug().Offset(-1).Count(&count)

	for _, projectPool := range projectPools {

		project := []model.Project{}

		projectPoolList := []int64(projectPool.ProjectList)

		db.Debug().Where("id IN ?", projectPoolList).Find(&project)

		projectPoolsResponse = append(projectPoolsResponse, project_pool_response_formatter.ProjectPoolResponseFormatter(projectPool, project))
	}

	return c.JSON(fiber.Map{
		"status":            "ok",
		"project_pool_list": projectPoolsResponse,
		"amount":            count,
	})
}

func ProjectPoolById(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	projectPool := model.ProjectPool{}

	db.Debug().Where("id = ?", id).Preload("TierList", "deleted_at IS NULL").First(&projectPool)

	project := []model.Project{}

	projectPoolList := []int64(projectPool.ProjectList)

	db.Debug().Where("id IN ?", projectPoolList).Order("created_at asc").Find(&project)

	projectPoolResponse := project_pool_response_formatter.ProjectPoolResponseFormatter(projectPool, project)

	return c.JSON(map[string]interface{}{
		"status": "ok",
		"result": projectPoolResponse,
	})
}
