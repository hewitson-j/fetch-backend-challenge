package handlers

import (
	"fetch-backend-challenge/models"
	"fetch-backend-challenge/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt(ctx *gin.Context){
	var receipt models.Receipt

	if err := ctx.ShouldBindJSON(&receipt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	fmt.Println("Processing receipt...")

	receiptId := uuid.New().String()

	storage.SaveReceipt(receiptId, receipt)
	log.Println("Saved " + receiptId)

	ctx.JSON(http.StatusOK, gin.H{"id":receiptId})
}

func GetPoints(ctx *gin.Context){
	receiptId := ctx.Param("id")

	receipt, exists := storage.GetReceipt(receiptId)

	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	points := storage.CalculatePoints(receipt)

	ctx.JSON(http.StatusOK, gin.H{"points": points})
}