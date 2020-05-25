package routers

import (
	"io"
	"net/http"
	"os"
	"github.com/maximp14/golangreact/db"
)

func DownloadBanner(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Must send id", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(writer, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(writer, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(writer, OpenFile)
	if err != nil {
		http.Error(writer, "Copy image error", http.StatusBadRequest)
	}

}
