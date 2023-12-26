package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/igorferrati/api-rest-golang/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
