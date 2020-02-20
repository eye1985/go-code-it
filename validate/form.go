package validate

import (
	"errors"
	"regexp"
)

func Username(username string) error {
	if len(username) < 3 {
		return errors.New("username must be 3 char or longer")
	}
	return nil
}

func Password(password string) error {
	if len(password) < 8 {
		return errors.New("password must be longer than 8 char")
	}

	return nil
}

func Email(email string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return errors.New("not a valid email address")
	}
	return nil
}

func CodeTitle(title string) error {
	return nil
}

func CodeType(title string) error {
	return nil
}
