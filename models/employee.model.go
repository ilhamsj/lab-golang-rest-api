package model

type Employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Response struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    []Employee `json:"data"`
	Version int        `json:"version"`
	Status  int        `json:"status"`
}
