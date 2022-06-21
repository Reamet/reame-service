package changer

import (
	"bsc-scan-data-service/database/model"
	"bytes"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BscProjectPayload struct {
	ReferredId                      *float64   `json:"id"`
	Contract                        string     `json:"contract"`
	ContractVersion                 int        `json:"contract_version"`
	OpenTime                        int        `json:"openTime"`
	CloseTime                       int        `json:"closeTime"`
	Decimals                        int        `json:"decimals"`
	State                           string     `json:"state"`
	IsPrivate                       bool       `json:"isPrivate"`
	Rate                            string     `json:"rate"`
	TotalCountWallet                int        `json:"totalCountWallet"`
	TotalCountUserParticipated      int        `json:"totalCountUserParticipated"`
	TotalFundParticipated           float64    `json:"totalFundParticipated"`
	MaxSingleParticipationAllocated int        `json:"maxSingleParticipationAllocated"`
	MaxTotalParticipationAllocated  string     `json:"maxTotalParticipationAllocated"`
	Description                     string     `json:"description"`
	Telegram                        *string    `json:"telegram"`
	ProjectTokenAddress             string     `json:"projectTokenAddress"`
	Logo                            string     `json:"logo"`
	Medium                          *string    `json:"medium"`
	Name                            string     `json:"name"`
	ProjectTokenSymbol              string     `json:"projectTokenSymbol"`
	TotalSupply                     string     `json:"total_supply"`
	Twitter                         *string    `json:"twitter"`
	Website                         string     `json:"website"`
	YourAllocationVisible           bool       `json:"yourAllocationVisible"`
	Detail                          *string    `json:"detail"`
	ProjectTokenContract            *string    `json:"projectTokenContract"`
	AthMultiplier                   *int       `json:"athMultiplier"`
	Symbol                          string     `json:"symbol"`
	Disabled                        bool       `json:"disabled"`
	Pancakeswap                     string     `json:"pancakeswap"`
	Start                           *time.Time `json:"start"`
	End                             *time.Time `json:"end"`
	Staking                         string     `json:"staking"`
	Allocation                      string     `json:"allocation"`
	Fcfs                            string     `json:"fcfs"`
	AllDay                          *bool      `json:"allDay"`
	TokenAddress                    string     `json:"tokenAddress"`
}

type ProjectPayload struct {
	ProjectSource string              `json:"projectSource"`
	ProjectList   []BscProjectPayload `json:"projectList"`
}

func ChangerJson(c *fiber.Ctx) error {
	bodyPayload := ProjectPayload{}

	if err := c.BodyParser(&bodyPayload); err != nil {
		return err
	}

	var projectDBPayloadTest []model.Project

	for _, project := range bodyPayload.ProjectList {

		var information bytes.Buffer

		information.WriteString("Opens " + strconv.Itoa(project.OpenTime) + "\r\n")
		information.WriteString("FCFS " + project.Fcfs + "\r\n")
		information.WriteString("Closes " + strconv.Itoa(project.CloseTime) + "\r\n")
		information.WriteString("Token Symbol " + project.Symbol + "\r\n")
		information.WriteString("Total Supply " + project.TotalSupply + "\r\n")

		databasePayload := model.Project{
			ReferredId:  project.ReferredId,
			Title:       project.Name,
			Description: project.Description,
			Logo:        project.Logo,
			Website:     project.Website,
			Information: information.String(),
			Telegram:    project.Telegram,
			Twitter:     project.Twitter,
			Discord:     nil,
			Email:       nil,
			Facebook:    nil,
			Instagram:   nil,
		}

		projectDBPayloadTest = append(projectDBPayloadTest, databasePayload)

	}

	return c.JSON(fiber.Map{
		"result": projectDBPayloadTest,
	})
}
