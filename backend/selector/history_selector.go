package selector

import (
	"time"

	"github.com/Oozaku/bug-bug-tracker/backend/database/models"
)

type HistoryApi struct {
	ActionDescription string
	ChangedTo         string
	UserUsername      string
	IssueID           uint
	CreatedAt         time.Time
}

func GetHistoryData(history models.History) HistoryApi {
	return HistoryApi{
		history.ActionDescription,
		history.ChangedTo,
		history.UserUsername,
		history.IssueID,
		history.CreatedAt,
	}
}
