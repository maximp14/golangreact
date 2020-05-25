package routers

import (
	"github.com/maximp14/golangreact/db"
	"net/http"
)

func DeleteTweet(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Must send the id", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweetFromDB(ID, IDUser)
	if err != nil {
		http.Error(writer, "Something went wrong deleting the tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}
