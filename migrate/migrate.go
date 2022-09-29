package main

import (
	"postgres/initializers"
	"postgres/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connecting_db()
}

func main() {
	initializers.DB.AutoMigrate(&models.Passenger{}, &models.Train{})
}
