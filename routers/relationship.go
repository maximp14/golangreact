package routers

import (
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
	"net/http"
)

func Relationship(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Must send id", http.StatusBadRequest)
		return
	}

	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID

	status, err := db.AddRelationship(t)
	if err != nil {
		http.Error(writer, "Relationship error"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(writer, "Error inserting relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
