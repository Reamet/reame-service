package project_pool

import (
	"bsc-scan-data-service/database/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProjectPoolTier struct {
	Tier        int `json:"tier"`
	TokenAmount int `json:"tokenAmount"`
}

type ProjectPoolPoll struct {
	Title string `json:"title"`
}

type ProjectPoolCreatePayload struct {
	Title                     string            `json:"title"`
	SubTitle                  string            `json:"subTitle"`
	Description               string            `json:"description"`
	Source                    string            `json:"source"`
	StartDate                 time.Time         `json:"startDate"`
	EndDate                   time.Time         `json:"endDate"`
	ProjectList               []int64           `json:"projectList"`
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
	PollList                  []ProjectPoolPoll `json:"pollList"`
}

func ProjectPoolCreate(c *fiber.Ctx, db *gorm.DB) error {
	bodyPayload := ProjectPoolCreatePayload{}

	currentTime := time.Now()

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	poolDatabasePayload := model.ProjectPool{
		Title:                     bodyPayload.Title,
		SubTitle:                  bodyPayload.SubTitle,
		Description:               bodyPayload.Description,
		Source:                    bodyPayload.Source,
		StartDate:                 bodyPayload.StartDate,
		EndDate:                   bodyPayload.EndDate,
		ProjectList:               bodyPayload.ProjectList,
		Term:                      bodyPayload.Term,
		InvestmentPeriod:          int(bodyPayload.InvestmentPeriod),
		WithdrawalDate:            bodyPayload.WithdrawalDate,
		StartVoteDate:             bodyPayload.StartVoteDate,
		EndVoteDate:               bodyPayload.EndVoteDate,
		GoalRaised:                int(bodyPayload.GoalRaised),
		GoalAllocation:            int(bodyPayload.GoalAllocation),
		BasicInvestmentSuggestion: int(bodyPayload.BasicInvestmentSuggestion),
		DepositFee:                int(bodyPayload.DepositFee),
		Ido:                       int(bodyPayload.Ido),
		Stake:                     int(bodyPayload.Stake),
		Status:                    bodyPayload.Status,
		UpdatedAt:                 currentTime,
		CreatedAt:                 currentTime,
	}

	err := db.Debug().Create(&poolDatabasePayload).Error

	if err != nil {
		return err
	}

	for _, tier := range bodyPayload.TierList {
		tierDatabasePayload := model.ProjectTier{
			PoolId:      poolDatabasePayload.ID,
			Tier:        tier.Tier,
			TokenAmount: tier.TokenAmount,
		}

		tierErr := db.Debug().Create(&tierDatabasePayload).Error

		if tierErr != nil {
			return tierErr
		}
	}

	if bodyPayload.PollList != nil {
		for _, poll := range bodyPayload.PollList {
			pollDatabasePayload := model.ProjectPoll{
				PoolId: poolDatabasePayload.ID,
				Title:  poll.Title,
			}

			pollErr := db.Debug().Create(&pollDatabasePayload).Error

			if pollErr != nil {
				return pollErr
			}
		}
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
