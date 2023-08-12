package database

import "github.com/Oozaku/bug-bug-tracker/backend/database/models"

func CreateIssue(issue *models.Issue) (*models.Issue, error) {
	result := DBConnection.Create(issue)

	if result.Error != nil {
		return &models.Issue{}, result.Error
	}

	return issue, nil
}
