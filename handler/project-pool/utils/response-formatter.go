package project_pool_response_formatter

import "bsc-scan-data-service/database/model"

func ProjectPoolResponseFormatter(projectPool model.ProjectPool, project []model.Project) model.ProjectPoolResponse {
	projectPoolResponse := model.ProjectPoolResponse{
		ID:                        projectPool.ID,
		Title:                     projectPool.Title,
		SubTitle:                  projectPool.SubTitle,
		Description:               projectPool.Description,
		Source:                    projectPool.Source,
		StartDate:                 projectPool.StartDate,
		EndDate:                   projectPool.EndDate,
		StartVoteDate:             projectPool.StartVoteDate,
		EndVoteDate:               projectPool.EndVoteDate,
		ProjectList:               project,
		Term:                      projectPool.Term,
		InvestmentPeriod:          projectPool.InvestmentPeriod,
		WithdrawalDate:            projectPool.WithdrawalDate,
		GoalRaised:                projectPool.GoalRaised,
		GoalAllocation:            projectPool.GoalAllocation,
		BasicInvestmentSuggestion: projectPool.BasicInvestmentSuggestion,
		DepositFee:                projectPool.DepositFee,
		Ido:                       projectPool.Ido,
		Stake:                     projectPool.Stake,
		Status:                    projectPool.Status,
		PoolAddress:               projectPool.PoolAddress,
		TierList:                  projectPool.TierList,
		PollList:                  projectPool.PollList,
		UpdatedAt:                 projectPool.UpdatedAt,
		CreatedAt:                 projectPool.CreatedAt,
		DeletedAt:                 projectPool.DeletedAt,
	}

	return projectPoolResponse
}
