package routes

import (
	"fetch-backend-challenge/handlers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Success!"})

		fmt.Println("Test route accessed.")
	})

	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetPoints)

	return r
}