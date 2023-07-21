package selector

import "github.com/Oozaku/bug-bug-tracker/backend/database/models"

type UserPublicApi struct {
	Username  string
	Name      string
	Histories []HistoryApi
}

func GetPublicUserData(user models.User) UserPublicApi {

	if user.Histories == nil {
		return UserPublicApi{
			user.Username,
			user.Name,
			nil,
		}
	}

	var histories []HistoryApi

	for _, history := range user.Histories {
		histories = append(histories, GetHistoryData(history))
	}

	return UserPublicApi{
		user.Username,
		user.Name,
		histories,
	}
}
