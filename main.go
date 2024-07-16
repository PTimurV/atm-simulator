package main

import (
	"log"
	"net/http"

	"atm-simulator/handlers"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
    r.HandleFunc("/accounts/{id}/deposit", handlers.Deposit).Methods("POST")
    r.HandleFunc("/accounts/{id}/withdraw", handlers.Withdraw).Methods("POST")
    r.HandleFunc("/accounts/{id}/balance", handlers.GetBalance).Methods("GET")

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
