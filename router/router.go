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

	reameServiceCollectionGroup := api.Group("/collection")
	reameServiceCollectionGroupHandler := handler.CollectionHandler{}
	reameServiceCollectionGroupHandler.Init(database.Database.DB)
	reameServiceCollectionGroup.Post("/create", reameServiceCollectionGroupHandler.Create)
	reameServiceCollectionGroup.Get("/lists", reameServiceCollectionGroupHandler.CollectionLists)
	reameServiceCollectionGroup.Get("/short_url", reameServiceCollectionGroupHandler.CollectionByShortUrl)
	reameServiceCollectionGroup.Get("/chain/collection_id_chain", reameServiceCollectionGroupHandler.CollectionByIdChain)
	reameServiceCollectionGroup.Post("/update/:id", reameServiceCollectionGroupHandler.Update)
	reameServiceCollectionGroup.Get("/:id", reameServiceCollectionGroupHandler.CollectionById)

	reameServiceMintGroup := api.Group("/mint")
	reameServiceMintGroupHandler := handler.MintHandler{}
	reameServiceMintGroupHandler.Init(database.Database.DB)
	reameServiceMintGroup.Post("/create", reameServiceMintGroupHandler.Create)
	reameServiceMintGroup.Post("/update/:id", reameServiceMintGroupHandler.Update)
	reameServiceMintGroup.Get("/lists", reameServiceMintGroupHandler.MintLists)
	reameServiceMintGroup.Get("/:id", reameServiceMintGroupHandler.GetMintById)

	reameServiceLaunchpadGroup := api.Group("/launchpad")
	reameServiceLaunchpadGroupHandler := handler.LaunchpadHandler{}
	reameServiceLaunchpadGroupHandler.Init(database.Database.DB)
	reameServiceLaunchpadGroup.Post("/create", reameServiceLaunchpadGroupHandler.Create)
	reameServiceLaunchpadGroup.Put("/update", reameServiceLaunchpadGroupHandler.Update)
	reameServiceLaunchpadGroup.Get("/", reameServiceLaunchpadGroupHandler.GetLaunchPadAll)
	reameServiceLaunchpadGroup.Get("/:id", reameServiceLaunchpadGroupHandler.GetLaunchPadById)
	reameServiceLaunchpadGroup.Get("/slug/:slug", reameServiceLaunchpadGroupHandler.GetLaunchPadBySlug)
}
