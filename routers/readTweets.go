package routers

import (
	"encoding/json"
	"github.com/maximp14/golangreact/db"
	"net/http"
	"strconv"
)

func ReadTweets(writer http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Must send id", http.StatusBadRequest)
		return
	}

	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(writer, "Must send page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		http.Error(writer, "Must send page with a value", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	resp, success := db.ReadTweets(ID, pag)
	if success == false {
		http.Error(writer, "Error reading tweets", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resp)
}
