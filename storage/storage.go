package storage

import (
	"fetch-backend-challenge/models"
	"math"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var (
	receipts = make(map[string]models.Receipt)
	mu sync.Mutex
)

func SaveReceipt(id string, r models.Receipt){
	mu.Lock()
	defer mu.Unlock()
	receipts[id] = r
}

func GetReceipt(id string)(models.Receipt, bool){
	mu.Lock()
	defer mu.Unlock()

	receipt, exists := receipts[id]

	return receipt, exists
}

func CalculatePoints(r models.Receipt) int {
	points := 0

	for _, char := range r.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	price, err := strconv.ParseFloat(r.Total, 32)
	if err == nil && math.Mod(price, 1) == 0 {
		points += 50
	}

	if err == nil && math.Mod(price, .25) == 0 {
		points += 25
	}

	items := r.Items
	points += (len(items) / 2) * 5

	for _, item := range items {
		shortDescription := strings.TrimSpace(item.ShortDescription)
		shortDescriptionLen := len(shortDescription)

		if shortDescriptionLen % 3 == 0 {
			itemPriceStr := item.Price
			itemPrice, err := strconv.ParseFloat(itemPriceStr, 64)

			if err == nil {
				itemPrice *= .2

				points += int(math.Ceil(itemPrice))
			}
		}
	}

	date, err2 := strconv.Atoi(r.PurchaseDate[len(r.PurchaseDate) - 1:]) 

	if err2 == nil && date % 2 != 0 {
		points += 6
	}

	// TODO: 10 points if purchased between 2 and 4 PM

	return points
}