package handler

import (
	"reame-service/handler/collection"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CollectionHandler struct {
	DB *gorm.DB
}

func (ph *CollectionHandler) Init(db *gorm.DB) {
	ph.DB = db
}

func (ph *CollectionHandler) Create(c *fiber.Ctx) error {
	return collection.Create(c, ph.DB)
}

func (ph *CollectionHandler) Update(c *fiber.Ctx) error {
	return collection.Update(c, ph.DB)
}

func (ph *CollectionHandler) CollectionLists(c *fiber.Ctx) error {
	return collection.CollectionLists(c, ph.DB)
}

func (ph *CollectionHandler) CollectionById(c *fiber.Ctx) error {
	return collection.CollectionById(c, ph.DB)
}

func (ph *CollectionHandler) CollectionByShortUrl(c *fiber.Ctx) error {
	return collection.CollectionByShortUrl(c, ph.DB)
}

func (ph *CollectionHandler) CollectionByIdChain(c *fiber.Ctx) error {
	return collection.CollectionByIdChain(c, ph.DB)
}

func (ph *CollectionHandler) CreateTrendingCollection(c *fiber.Ctx) error {
	return collection.CreateTrendingCollection(c, ph.DB)
}

func (ph *CollectionHandler) UpdateTrendingCollection(c *fiber.Ctx) error {
	return collection.UpdateTrendingCollection(c, ph.DB)
}

func (ph *CollectionHandler) GetAllTrendingCollection(c *fiber.Ctx) error {
	return collection.GetAllTrendingCollection(c, ph.DB)
}

func (ph *CollectionHandler) CreateFeaturedCollection(c *fiber.Ctx) error {
	return collection.CreateFeaturedCollection(c, ph.DB)
}

func (ph *CollectionHandler) UpdateFeaturedCollection(c *fiber.Ctx) error {
	return collection.UpdateFeaturedCollection(c, ph.DB)
}

func (ph *CollectionHandler) GetAllFeaturedCollection(c *fiber.Ctx) error {
	return collection.GetAllFeaturedCollection(c, ph.DB)
}
