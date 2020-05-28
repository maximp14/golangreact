package routers

import (
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
	"net/http"
)

func RemoveRelationship(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Must send id", http.StatusBadRequest)
		return
	}

	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID

	status, err := db.RemoveRelationship(t)
	if err != nil {
		http.Error(writer, "Remove relationship error"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(writer, "Error deleting relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
