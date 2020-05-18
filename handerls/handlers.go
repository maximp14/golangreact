package handerls

import (
	"github.com/gorilla/mux"
	"github.com/maximp14/golangreact/middlew"
	"github.com/maximp14/golangreact/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	router := mux.NewRouter()

	//	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT,handler))


}
