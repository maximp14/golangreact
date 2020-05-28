package routers

import (
	"encoding/json"
	"net/http"

	"github.com/maximp14/golangreact/db"
	"strconv"
)

func PersonalTL(writer http.ResponseWriter, request *http.Request) {
	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(writer, "Must send page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		http.Error(writer, "Must send page > 0", http.StatusBadRequest)
		return
	}

	resp, status := db.PersonalTL(IDUser, page)
	if status == false {
		http.Error(writer, "Error reading tweets", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}
