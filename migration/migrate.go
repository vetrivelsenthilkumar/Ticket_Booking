package migration

import (
	"Ticket_Booking_App/initializers"
	"Ticket_Booking_App/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connecting_db()
}

func MigrateTable() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Train{})
}
