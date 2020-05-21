package routers

import (
	"encoding/json"
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/jwt"
	"github.com/maximp14/golangreact/models"
	"net/http"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-type", "application/json")

	var t models.User

	err:=json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "User or Password invalid"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(writer, "Email required"+err.Error(), 400)
		return
	}

	document, exist := db.LoginTry(t.Email, t.Password)

	if exist == false {
		http.Error(writer, "User or Password invalid", 400)
		return
	}

	jwtKey, err := jwt.GeneratedJWT(document)
	if err != nil {
		http.Error(writer, "Token error"+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(writer, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})


}