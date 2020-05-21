package db

import (
	"github.com/maximp14/golangreact/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginTry(email string, password string) (models.User, bool) {
	user, found, _ := UserExist(email)

	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
