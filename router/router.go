package router

import (
	"fmt"
	"reame-service/database"
	"reame-service/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const VERSION = "v1"

type socketMessage struct {
	Title string `json:"title"`
}

func SetRouter(app *fiber.App) {
	version := app.Group("/" + VERSION)
	api := version.Group("/", logger.New())

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("Hello! This is Seedtopia RESTful API. 👏 %s", VERSION)})
	})

	reameServiceGroup := api.Group("/collection")
	reameServiceGroupHandler := handler.ProjectHandler{}
	reameServiceGroupHandler.Init(database.Database.DB)
	reameServiceGroup.Post("/create", reameServiceGroupHandler.Create)

}
