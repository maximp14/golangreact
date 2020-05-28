package routers

import (
	"encoding/json"
	"github.com/maximp14/golangreact/db"
	"net/http"
	"strconv"
)

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	typeUser := request.URL.Query().Get("type")
	page := request.URL.Query().Get("page")
	search := request.URL.Query().Get("search")

	pageTmp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(writer, "Must send page > 0", http.StatusBadRequest)
		return
	}
	pag := int64(pageTmp)

	result, status := db.TimeLine(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(writer, "Error reading users", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)

}
