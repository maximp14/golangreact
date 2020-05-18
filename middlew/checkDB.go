package middlew

import (
	"net/http"
	"github.com/maximp14/golangreact/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if db.CheckConnection() == false {
			http.Error(writer, "DB Connection lost", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
