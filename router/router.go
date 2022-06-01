package router

import (
	"bsc-scan-data-service/database"
	"bsc-scan-data-service/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const VERSION = "v1"

type socketMessage struct {
	Title   string `json:"title"`
}

func SetRouter(app *fiber.App) {
	version := app.Group("/" + VERSION)
	api := version.Group("/", logger.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("Hello! This is Notification RESTful API. 👏 %s", VERSION)})
	})

	bscProjectGroup := api.Group("/")
	bscProjectGroupHandler := handler.ProjectHandler{}
	bscProjectGroupHandler.Init(database.Database.DB)
	bscProjectGroup.Post("/create-projects", bscProjectGroupHandler.ProjectCreate)
}