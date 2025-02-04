package main

import (
	"fetch-backend-challenge/routes"
)

func main(){
	r := routes.SetupRouter()

	r.Run(":8080")
}