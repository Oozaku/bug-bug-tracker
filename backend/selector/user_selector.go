package selector

import "github.com/Oozaku/bug-bug-tracker/backend/database/models"

type UserPublicApi struct {
	Username  string
	Name      string
	Histories []HistoryApi
}

type UserPrivateApi struct {
	Username  string
	Name      string
	Email     string
	Role      string
	Histories []HistoryApi
}

func GetPrivateUserData(user models.User) UserPrivateApi {

	if user.Histories == nil {
		return UserPrivateApi{
			user.Username,
			user.Name,
			user.Email,
			user.Role,
			nil,
		}
	}

	var histories []HistoryApi

	for _, history := range user.Histories {
		histories = append(histories, GetHistoryData(history))
	}

	return UserPrivateApi{
		user.Username,
		user.Name,
		user.Email,
		user.Role,
		histories,
	}
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
