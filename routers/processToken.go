package routers

import (
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
)

var Email string

var IDUser string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("ExtraDosE2")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil{
		_, found, _ := db.UserExist(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}