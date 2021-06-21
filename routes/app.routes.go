package routes

import (
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/employee", controller.AllEmployee).Methods("GET")
	router.HandleFunc("/employee/store", controller.InsertEmployee).Methods("POST")
}
