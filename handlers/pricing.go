package handlers

import (
	"encoding/json"
	"net/http"

	"canteen-pricing/models"
	"canteen-pricing/pricing"
)

func CalculatePrice(w http.ResponseWriter, r *http.Request) {

	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var request models.PricingRequest

	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create closure
	discounter := pricing.CreateDiscounter(request.Discount)

	// Calculate values
	finalPrice := discounter(request.Price)
	discountAmount := request.Price - finalPrice

	response := models.PricingResponse{
		ItemName:        request.ItemName,
		OriginalPrice:   request.Price,
		DiscountPercent: request.Discount,
		DiscountAmount:  discountAmount,
		FinalPrice:      finalPrice,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}