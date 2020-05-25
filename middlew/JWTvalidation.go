package middlew

import (
	"net/http"
	"github.com/maximp14/golangreact/routers"
)

func JWTValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.ProcessToken(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(writer, "Token error "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
