package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPwd string, plainPass []byte) (bool, error) {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		return false, err
	}

	return true, nil
}
