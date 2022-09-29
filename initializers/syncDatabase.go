package initializers

import "postgres/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Passenger{}, &models.Train{})
}
