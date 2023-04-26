package router

import (
	"fmt"
	"reame-service/database"
	"reame-service/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const VERSION = "v1"

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
	reameServiceCollectionGroup.Get("/trending", reameServiceCollectionGroupHandler.GetAllTrendingCollection)
	reameServiceCollectionGroup.Post("/trending/create", reameServiceCollectionGroupHandler.CreateTrendingCollection)
	reameServiceCollectionGroup.Put("/trending/update/:id", reameServiceCollectionGroupHandler.UpdateTrendingCollection)
	reameServiceCollectionGroup.Get("/featured", reameServiceCollectionGroupHandler.GetAllFeaturedCollection)
	reameServiceCollectionGroup.Post("/featured/create", reameServiceCollectionGroupHandler.CreateFeaturedCollection)
	reameServiceCollectionGroup.Put("/featured/update/:id", reameServiceCollectionGroupHandler.UpdateFeaturedCollection)
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

	reameHomeGroup := api.Group("/home")
	reameHomeGroupHandler := handler.HomeHandler{
		DB: database.Database.DB,
	}
	reameHomeGroupHandler.Init(database.Database.DB)
	reameHomeGroup.Post("/create", reameHomeGroupHandler.Create)
	reameHomeGroup.Put("/update/:id", reameHomeGroupHandler.Update)
	reameHomeGroup.Get("/", reameHomeGroupHandler.GetAllHome)
	reameHomeGroup.Get("/:id", reameHomeGroupHandler.GetHomeById)

	excollectionHandler := handler.ExCollectionHandler{
		DB: database.Database.DB,
	}
	excollection := api.Group("/excollection")
	excollection.Get("/", excollectionHandler.GetCollections)
	excollection.Get("/:ref", excollectionHandler.GetCollectionSingleDetail)
	excollection.Get("/exist", excollectionHandler.GetIsExist)
	excollection.Post("/", excollectionHandler.PostCreateNewCollectionDetail)
	excollection.Put("/", excollectionHandler.PutUpdateCollectionDetail)

}
