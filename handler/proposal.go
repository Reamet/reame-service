package handler

import (
	proposal "bsc-scan-data-service/handler/proposal"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProposalHandler struct {
	DB *gorm.DB
}

func (pph *ProposalHandler) Init(db *gorm.DB) {
	pph.DB = db
}

func (pph *ProposalHandler) ProposalCreate(c *fiber.Ctx) error {
	return proposal.ProposalCreate(c, pph.DB)
}

func (pph *ProposalHandler) ProposalUpdate(c *fiber.Ctx) error {
	return proposal.ProposalUpdate(c, pph.DB)
}

func (pph *ProposalHandler) ProposalById(c *fiber.Ctx) error {
	return proposal.ProposalById(c, pph.DB)
}

func (pph *ProposalHandler) ProposalLists(c *fiber.Ctx) error {
	return proposal.ProposalLists(c, pph.DB)
}
