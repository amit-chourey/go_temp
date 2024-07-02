package api

import (
	"encoding/json"
	"net/http"
)

// coint Balance params
type CoinBalanceParams struct{
	username string
}

type CoinBalanceResponse struct {
	// success code, usually 200
	code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error code
	code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)