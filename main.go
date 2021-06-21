package main

import (
	"fmt"
	"myapp/controller"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/employee", controller.AllEmployee).Methods("GET")
	router.HandleFunc("/employee/store", controller.InsertEmployee).Methods("POST")

	fmt.Println("Connected to port 3001")
	http.ListenAndServe(":3001", router)
}
