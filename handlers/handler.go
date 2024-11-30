package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/utils"
	"sync"
)

var (
	receiptStore = make(map[string]int) // Receipt id to points map
	mutex        = &sync.Mutex{}        // Mutex for thread safety
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Calculate points
	points := utils.CalculatePoints(receipt)

	// Generate a unique receipt ID (UUID format)
	receiptID := uuid.New().String()

	// Store receipt and points
	mutex.Lock()
	receiptStore[receiptID] = points
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": receiptID})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid receipt id", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	points, exists := receiptStore[id]
	mutex.Unlock()

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
