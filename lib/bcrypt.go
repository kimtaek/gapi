package lib

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return err.Error()
	}
	return string(encryptPassword)
}

func ComparePassword(old string, new string) bool {
	return bcrypt.CompareHashAndPassword([]byte(old), []byte(new)) == nil
}
