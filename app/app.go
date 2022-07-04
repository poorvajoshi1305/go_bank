package main

import (
	"net/http"
)

func Start() {

	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/getAllCustomers", getAllCustomers)

	//starting server
	http.ListenAndServe("localhost:5000", nil)
}
