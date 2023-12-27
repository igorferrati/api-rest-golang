package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/igorferrati/api-rest-golang/controllers"
	"github.com/igorferrati/api-rest-golang/middleware"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContenTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
