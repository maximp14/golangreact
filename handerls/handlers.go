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

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckDB(middlew.JWTValidation(routers.LookProfile))).Methods("GET")
	router.HandleFunc("/editProfile", middlew.CheckDB(middlew.JWTValidation(routers.EditProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.JWTValidation(routers.TweetPersist))).Methods("POST")
	router.HandleFunc("/readTweets", middlew.CheckDB(middlew.JWTValidation(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.JWTValidation(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.JWTValidation(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/downloadAvatar", middlew.CheckDB(routers.DownloadAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.JWTValidation(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/downloadBanner", middlew.CheckDB(routers.DownloadBanner)).Methods("GET")

	router.HandleFunc("/addRelationship", middlew.CheckDB(middlew.JWTValidation(routers.Relationship))).Methods("POST")
	router.HandleFunc("/removeRelationship", middlew.CheckDB(middlew.JWTValidation(routers.RemoveRelationship))).Methods("DELETE")
	router.HandleFunc("/haveRelationship", middlew.CheckDB(middlew.JWTValidation(routers.HaveRelationship))).Methods("GET")

	router.HandleFunc("/listUsers", middlew.CheckDB(middlew.JWTValidation(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/personalTL", middlew.CheckDB(middlew.JWTValidation(routers.PersonalTL))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
