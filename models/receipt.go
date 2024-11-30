package models

type Receipt struct {
	Retailer     string `json:"retailer"`     // retailer
	PurchaseDate string `json:"purchaseDate"` // purchase date YYYY-MM-DD
	PurchaseTime string `json:"purchaseTime"` // purchase time HH:MM
	Items        []Item `json:"items"`        // items purchased in the receipt
	Total        string `json:"total"`        // total amount purchased in the receipt
}

type Item struct {
	ShortDescription string `json:"shortDescription"` // the short description of the item
	Price            string `json:"price"`            // the price of the item
}
