package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	CustId  int    `json:"cust_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Poorva")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{
		{1, "Poorva", "Navi-Mumbai"},
		{2, "pooja", "Pune"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)

}
