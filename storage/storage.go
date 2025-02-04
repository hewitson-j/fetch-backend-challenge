package storage

import (
	"fetch-backend-challenge/models"
	"sync"
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
	return 100
}