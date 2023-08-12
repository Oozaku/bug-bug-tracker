package validator

import (
	"errors"
	"regexp"
)

func ValidateUsersUsername(username string) error {
	if len(username) < 4 || len(username) > 15 {
		return errors.New("Username's size is out of range [5, 15]")
	}

	usernameIsValid, err := regexp.MatchString("^[0-9A-z]+$", username)

	if err != nil {
		return errors.New("Unknown error while validating username")
	}

	if !usernameIsValid {
		return errors.New("Username is invalid")
	}

	return nil
}
