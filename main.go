package main

import (
	"Ticket_Booking_App/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connecting_db()
}

func main() {
	r := gin.Default()
	r.GET("/")
	r.Run()
}
