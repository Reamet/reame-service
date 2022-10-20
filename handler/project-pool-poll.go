package handler

import (
	pool_poll "bsc-scan-data-service/handler/pool-poll"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PollHandler struct {
	DB *gorm.DB
}

func (pph *PollHandler) Init(db *gorm.DB) {
	pph.DB = db
}

func (pph *PollHandler) PollCreate(c *fiber.Ctx) error {
	return pool_poll.PollCreate(c, pph.DB)
}

func (pph *PollHandler) PollListByPoolId(c *fiber.Ctx) error {
	return pool_poll.PollListByPoolId(c, pph.DB)
}

func (pph *PollHandler) PollResultByPollId(c *fiber.Ctx) error {
	return pool_poll.PollResultByPollId(c, pph.DB)
}
