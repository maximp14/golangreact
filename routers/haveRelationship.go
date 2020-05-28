package routers

import (
	"encoding/json"
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
	"net/http"
)

func HaveRelationship(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")

	var t models.Relationship
	t.UserID = IDUser
	t.UserRelationshipID = ID

	var resp models.ResponseHaveRelationship

	status, err := db.HaveRelationship(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}
