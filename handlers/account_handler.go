package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"atm-simulator/services"

	"github.com/gorilla/mux"
)

var accountService = services.NewAccountService()

type CreateAccountRequest struct {
    Login string `json:"login"`
}

type CreateAccountResponse struct {
    ID string `json:"id"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
    var req CreateAccountRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    account, err := accountService.CreateAccount(req.Login)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    res := CreateAccountResponse{ID: account.ID}
    json.NewEncoder(w).Encode(res)
}

func Deposit(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
    if err != nil {
        http.Error(w, "Invalid amount", http.StatusBadRequest)
        return
    }
    err = accountService.Deposit(id, amount)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
    if err != nil {
        http.Error(w, "Invalid amount", http.StatusBadRequest)
        return
    }
    err = accountService.Withdraw(id, amount)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    balance, err := accountService.GetBalance(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}
