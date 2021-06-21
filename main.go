package main

import (
	"fmt"
	"myapp/controllers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/employee", controllers.EmployeeList).Methods("GET")
	router.HandleFunc("/employee/store", controllers.EmployeeStore).Methods("POST")

	fmt.Println("Connected to port 3001")
	http.ListenAndServe(":3001", router)
}
