package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/config"
	model "myapp/models"
	"net/http"
)

func EmployeeList(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	var response model.Response
	var arrEmployee []model.Employee

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, city FROM employee")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.City)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrEmployee = append(arrEmployee, employee)
		}
	}

	response.Success = true
	response.Message = "Data successfully retrieved"
	response.Data = arrEmployee
	response.Status = 200

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func EmployeeStore(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	city := r.FormValue("city")

	_, err = db.Exec("INSERT INTO employee(id, name, city) VALUES(?, ?, ?)", id, name, city)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	response.Success = true
	response.Data = []model.Employee{
		{
			Id:   id,
			Name: name,
			City: city,
		},
	}

	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
