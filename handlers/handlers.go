package handlers

import (
	"fetch-backend-challenge/models"
	"fetch-backend-challenge/storage"
	"fmt"
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

	ctx.JSON(http.StatusOK, gin.H{"id":receiptId})
}