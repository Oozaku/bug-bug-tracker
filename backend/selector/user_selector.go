package selector

import "github.com/Oozaku/bug-bug-tracker/backend/database/models"

type UserPublicApi struct {
	Username string
	Name     string
}

type UserPrivateApi struct {
	Username string
	Name     string
	Email    string
	Role     string
}

func GetPrivateUserData(user models.User) UserPrivateApi {

	return UserPrivateApi{
		user.Username,
		user.Name,
		user.Email,
		user.Role,
	}
}

func GetPublicUserData(user models.User) UserPublicApi {

	return UserPublicApi{
		user.Username,
		user.Name,
	}
}
