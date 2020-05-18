package routers

import (
	"encoding/json"
	"net/http"
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	var u models.User
	err := json.NewDecoder(request.Body).Decode(&u)
	if err != nil {
		http.Error(writer, "Data error"+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(writer, "Email required", 400)
	}
	if len(u.Password) < 6 {
		http.Error(writer, "Password must have at least 6 characters", 400)
	}

	_, found, _ := db.UserExist(u.Email)

	if found {
		http.Error(writer, "The email already exist in our data base", 400)
		return
	}

	_, status, err := db.InsertUser(u)
	if err != nil{
		http.Error(writer, "User register error"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(writer, "Cannot insert user registry ", 400)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}