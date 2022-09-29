package main

import (
	"Ticket_Booking_App/initializers"
	"Ticket_Booking_App/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connecting_db()
}

func main() {
	initializers.DB.AutoMigrate(&models.Passenger{}, &models.Train{})
}
