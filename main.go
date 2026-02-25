package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", whitelist)
	fmt.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type BookingParams struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Guests int    `json:"guests"`
	Date   string `json:"date"`
	Time   string `json:"time"`
}

func whitelist(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		methodOptionCORS(w)
		return
	}

	setCORS(w)
}

func methodOptionCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")
	w.WriteHeader(http.StatusNoContent)
}

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}
