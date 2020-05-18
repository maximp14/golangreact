package handerls

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/maximp14/golangreact/middlew"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT,handler))


}
