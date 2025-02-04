package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Success!"})

		fmt.Println("Test route accessed.")
	})

	r.Run(":8080")
}