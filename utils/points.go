package utils

import (
	"math"
	"receipt-processor-challenge/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	alphanumeric := regexp.MustCompile("[^a-zA-Z0-9]+")
	points += len(alphanumeric.ReplaceAllString(receipt.Retailer, ""))

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	total, err1 := strconv.ParseFloat(receipt.Total, 64)
	if err1 == nil && total == float64(int(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if err1 == nil && int(total*100)%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Item description length bonus
	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: Date bonus
	// 6 points if the day in the purchase date is odd.
	purchaseDate, err2 := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err2 == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: Time bonus
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err3 := time.Parse("15:04", receipt.PurchaseTime)
	if err3 == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points

}
