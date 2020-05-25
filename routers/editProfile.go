package routers

import (
	"encoding/json"
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
	"net/http"
)

func EditProfile(writer http.ResponseWriter, request *http.Request) {

	var t models.User

	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "Wrong information"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = db.EditUser(t, IDUser)
	if err != nil {
		http.Error(writer, "Something went wrong editing the profile data"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(writer, "Profile data not edited", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)

}
