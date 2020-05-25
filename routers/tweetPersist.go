package routers

import (
	"encoding/json"
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
	"net/http"
	"time"
)

func TweetPersist(writer http.ResponseWriter, request *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(request.Body).Decode(&message)

	data := models.TweetPersist{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.AddTweet(data)
	if err != nil {
		http.Error(writer, "Error adding the tweet, try again"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(writer, "Something went wrong adding the tweet, try again", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
