package routers

import (
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadBanner(writer http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(writer, "Image upload error"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(writer, "Something went wrong copying the image"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = db.EditUser(user, IDUser)
	if err != nil || status == false {
		http.Error(writer, "Error persisting the image in database"+err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

}
