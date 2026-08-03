package models

// Request received from the frontend
type PricingRequest struct {
	ItemName string  `json:"itemName"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
}

// Response sent back to the frontend
type PricingResponse struct {
	ItemName         string  `json:"itemName"`
	OriginalPrice    float64 `json:"originalPrice"`
	DiscountPercent  float64 `json:"discountPercent"`
	DiscountAmount   float64 `json:"discountAmount"`
	FinalPrice       float64 `json:"finalPrice"`
}