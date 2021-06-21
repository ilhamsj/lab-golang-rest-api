package routes

import (
	"myapp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/employee", controllers.EmployeeList).Methods("GET")
	router.HandleFunc("/employee/store", controllers.EmployeeStore).Methods("POST")
}
