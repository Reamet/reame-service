package project_pool

import (
	"bsc-scan-data-service/database/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ProjectPoolUpdatePayload struct {
	ID                        int               `json:"id"`
	Title                     string            `json:"title"`
	SubTitle                  string            `json:"subTitle"`
	Description               string            `json:"description"`
	Source                    string            `json:"source"`
	StartDate                 time.Time         `json:"startDate"`
	EndDate                   time.Time         `json:"endDate"`
	ProjectList               pq.Int64Array     `json:"projectList"`
	Term                      string            `json:"term"`
	InvestmentPeriod          int               `json:"investmentPeriod"`
	WithdrawalDate            time.Time         `json:"withdrawalDate"`
	StartVoteDate             time.Time         `json:"startVoteDate"`
	EndVoteDate               time.Time         `json:"endVoteDate"`
	GoalRaised                int               `json:"goalRaised"`
	GoalAllocation            int               `json:"goalAllocation"`
	BasicInvestmentSuggestion int               `json:"basicInvestmentSuggestion"`
	DepositFee                int               `json:"depositFee"`
	Ido                       int               `json:"ido"`
	Stake                     int               `json:"stake"`
	Status                    string            `json:"status"`
	TierList                  []ProjectPoolTier `json:"tierList"`
}

func ProjectPoolUpdate(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := ProjectPoolUpdatePayload{}

	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	databasePayload := map[string]interface{}{
		"title":                       bodyPayload.Title,
		"sub_title":                   bodyPayload.SubTitle,
		"description":                 bodyPayload.Description,
		"source":                      bodyPayload.Source,
		"start_date":                  bodyPayload.StartDate,
		"end_date":                    bodyPayload.EndDate,
		"project_list":                bodyPayload.ProjectList,
		"updated_at":                  currentTime,
		"term":                        bodyPayload.Term,
		"investment_period":           bodyPayload.InvestmentPeriod,
		"withdrawal_date":             bodyPayload.WithdrawalDate,
		"start_vote_date":             bodyPayload.StartVoteDate,
		"end_vote_date":               bodyPayload.EndVoteDate,
		"goal_raised":                 bodyPayload.GoalRaised,
		"goal_allocation":             bodyPayload.GoalAllocation,
		"basic_investment_suggestion": bodyPayload.BasicInvestmentSuggestion,
		"deposit_fee":                 bodyPayload.DepositFee,
		"ido":                         bodyPayload.Ido,
		"stake":                       bodyPayload.Stake,
		"status":                      bodyPayload.Status,
	}

	pool := model.ProjectPool{}

	poolResult := db.Debug().Where("ID = ?", bodyPayload.ID).First(&pool).Updates(&databasePayload)

	db.Debug().Model(&model.ProjectTier{}).Where("pool_id = ?", bodyPayload.ID).Updates(map[string]interface{}{
		"deleted_at": currentTime,
	})

	for _, tier := range bodyPayload.TierList {
		tierDatabasePayload := model.ProjectTier{
			PoolId:      bodyPayload.ID,
			Tier:        tier.Tier,
			TokenAmount: tier.TokenAmount,
		}

		tierErr := db.Debug().Create(&tierDatabasePayload).Error

		if tierErr != nil {
			return tierErr
		}
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": fmt.Sprintf("Row Affected By : %s row", strconv.FormatInt(poolResult.RowsAffected, 10)),
	})
}
