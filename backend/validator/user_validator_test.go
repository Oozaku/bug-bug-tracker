package validator

import "testing"

func TestFailOnEmptyUsername(t *testing.T) {
	emptyUsername := ""

	if err := ValidateUsersUsername(emptyUsername); err == nil {
		t.Errorf("Empty username should be invalidated")
	}
}

func TestFailOnLongUsername(t *testing.T) {
	longUsername := "sixteennnnnnnnnn"

	if err := ValidateUsersUsername(longUsername); err == nil {
		t.Errorf("Username with size bigger or equal to 16 should be invalidated")
	}
}

func TestFailOnShortUsername(t *testing.T) {
	shortUsername := "four"

	if err := ValidateUsersUsername(shortUsername); err == nil {
		t.Errorf("Username with length smaller than 5 should be invalidated")
	}
}

func TestFailOnUsernameWithSpace(t *testing.T) {
	usernameWithSpace := "space space"

	if err := ValidateUsersUsername(usernameWithSpace); err == nil {
		t.Errorf("Username with whitespace should be invalidated")
	}
}

func TestFailOnUsernameWithSpecialCharacters(t *testing.T) {
	usernamesWithSpecialCharacters := []string{
		"haha%%%",
		"hahaááá",
		"óáðßœfg",
		";;;;;;;",
		"\"\"\"\"\"",
		"=====",
	}

	for _, test := range usernamesWithSpecialCharacters {
		if err := ValidateUsersUsername(test); err == nil {
			t.Errorf("Username %q should have failed, because it has special characters", test)
		}
	}
}

func TestSucceedOnValidUsername(t *testing.T) {
	validUsernames := []string{
		"FIVEE",
		"fivee",
		"five5",
		"fifteeeeeeeeeen",
		"fifteen15151515",
	}

	for _, test := range validUsernames {
		if err := ValidateUsersUsername(test); err != nil {
			t.Errorf("Username %q should have succeed: error found: %q", test, err)
		}
	}
}
