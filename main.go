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

func whitelist(w http.ResponseWriter, r *http.Request) {

}
