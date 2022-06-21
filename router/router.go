package router

import (
	"bsc-scan-data-service/database"
	"bsc-scan-data-service/handler"
	"bsc-scan-data-service/handler/changer"
	"fmt"

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

	bscProjectGroup := api.Group("/")
	bscProjectGroupHandler := handler.ProjectHandler{}
	bscProjectGroupHandler.Init(database.Database.DB)
	bscProjectGroup.Post("/create-and-update-projects", bscProjectGroupHandler.ProjectCreate)
	bscProjectGroup.Get("/project/list-project", bscProjectGroupHandler.ProjectLists)
	bscProjectGroup.Get("/project/:id", bscProjectGroupHandler.ProjectById)

	poolProjectGroup := api.Group("/pool")
	poolProjectGroupHandler := handler.ProjectPoolHandler{}
	poolProjectGroupHandler.Init(database.Database.DB)
	poolProjectGroup.Post("/create-pool", poolProjectGroupHandler.ProjectPoolCreate)
	poolProjectGroup.Get("/list-pools", poolProjectGroupHandler.ProjectPoolList)
	poolProjectGroup.Post("/update-pool", poolProjectGroupHandler.ProjectPoolUpdate)
	poolProjectGroup.Get("/update-pool", poolProjectGroupHandler.ProjectPoolUpdate)
	poolProjectGroup.Get("/:id", poolProjectGroupHandler.ProjectPoolById)

	// This route use to convert payload from BSCPad to new one to use to insert in the new ways
	jsonChanger := api.Group("/json-changer")
	jsonChanger.Post("/json-bscpad-changer-v1", changer.ChangerJson)

}
