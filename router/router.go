package router

import (
	"fmt"
	"reame-service/database"
	"reame-service/handler"
	featuredlaunchpad "reame-service/handler/featuredLaunchpad"

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

	excollection := api.Group("/collection")
	excollectionHandler := handler.CollectionHandler{
		DB: database.Database.DB,
	}
	excollection.Get("/lists", excollectionHandler.GetCollections)
	excollection.Get("/:ref", excollectionHandler.GetCollectionSingleDetail)
	excollection.Get("/exist", excollectionHandler.GetIsExist)
	excollection.Post("/create", excollectionHandler.PostCreateNewCollectionDetail)
	excollection.Put("/update/:id", excollectionHandler.PutUpdateCollectionDetail)
	excollection.Patch("/:id/toggle-display", excollectionHandler.PatchUpdateToggleDisplay)

	excollection.Get("/trending/all", excollectionHandler.GetAllTrendingCollection)
	excollection.Post("/trending/create", excollectionHandler.CreateTrendingCollection)
	excollection.Put("/trending/update/:id", excollectionHandler.UpdateTrendingCollection)

	excollection.Get("/featured/all", excollectionHandler.GetAllFeaturedCollection)
	excollection.Post("/featured/create", excollectionHandler.CreateFeaturedCollection)
	excollection.Put("/featured/update/:id", excollectionHandler.UpdateFeaturedCollection)

	// owner route group
	ownerHandler := handler.OwnerHandler{
		DB: database.Database.DB,
	}
	owner := api.Group("/owner")
	owner.Get("/", ownerHandler.GetOwnerAll)
	owner.Get("/:ownerId", ownerHandler.GetOwnerById)
	owner.Post("/", ownerHandler.CreateOwner)
	owner.Put("/:ownerId", ownerHandler.UpdateOwner)

	// owner route group
	branchHandler := handler.BranchHandler{
		DB: database.Database.DB,
	}
	branch := api.Group("/branch")
	branch.Get("/", branchHandler.GetBranchAll)
	branch.Get("/:branchId", branchHandler.GetBranchById)
	branch.Get("/owner/:ownerId", branchHandler.GetBranchByOwnerId)
	branch.Post("/", branchHandler.CreateBranch)
	branch.Put("/:branchId", branchHandler.UpdateBranch)

	/* Featured Launchpad */
	featuredLaunchpad := api.Group("/featured")
	featuredLaunchpadHandler := featuredlaunchpad.FeaturedLaunchpadHandler{
		DB: database.Database.DB,
	}

	featuredLaunchpad.Get("/launchpad/all", featuredLaunchpadHandler.GetAllFeaturedLaunchpad)
	featuredLaunchpad.Post("/launchpad/create", featuredLaunchpadHandler.CreateFeaturedLaunchpad)
	featuredLaunchpad.Put("/launchpad/update/:id", featuredLaunchpadHandler.UpdateFeaturedLaunchpad)

}
