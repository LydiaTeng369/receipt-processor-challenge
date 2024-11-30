package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"receipt-processor-challenge/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
