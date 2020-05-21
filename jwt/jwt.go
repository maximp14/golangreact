package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maximp14/golangreact/models"
	"time"
)

func GeneratedJWT(t models.User) (string, error) {
	myKey := []byte("ExtraDosE2")

	payload := jwt.MapClaims{
		"email":      t.Email,
		"name":       t.Name,
		"last_name":  t.LastName,
		"birth_date": t.BirthDate,
		"biography":  t.Bio,
		"location":   t.Location,
		"website":    t.WebSite,
		"_id":        t.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
