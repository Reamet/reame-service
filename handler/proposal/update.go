package proposal

import (
	"bsc-scan-data-service/database/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProposalUpdatePayload struct {
	PoolAddress   string    `json:"poolAddress"`
	PollId        string    `json:"pollId"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	StartVoteDate time.Time `json:"startVoteDate"`
	EndVoteDate   time.Time `json:"endVoteDate"`
}

func ProposalUpdate(c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	currentTime := time.Now()

	bodyPayload := ProposalUpdatePayload{}
	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	databasePayload := map[string]interface{}{
		"title":           bodyPayload.Title,
		"description":     bodyPayload.Description,
		"pool_address":    bodyPayload.PoolAddress,
		"poll_id":         bodyPayload.PollId,
		"updated_at":      currentTime,
		"status":          bodyPayload.Status,
		"start_vote_date": bodyPayload.StartVoteDate,
		"end_vote_date":   bodyPayload.EndVoteDate,
	}

	proposalModel := model.Proposal{}

	poolResult := db.Debug().Where("ID = ?", id).First(&proposalModel).Updates(&databasePayload)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": fmt.Sprintf("Row Affected By : %s row", strconv.FormatInt(poolResult.RowsAffected, 10)),
	})
}
