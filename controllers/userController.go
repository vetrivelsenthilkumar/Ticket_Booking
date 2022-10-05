package controllers

import (
	"Ticket_Booking_App/initializers"
	"Ticket_Booking_App/models"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var passenger_detail struct {
		Name     string
		Age      int
		Email    string
		Password string
	}

	if c.Bind(&passenger_detail) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to register the details",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(passenger_detail.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to hash password",
		})
	}

	user := models.Passenger{Name: passenger_detail.Name, Age: passenger_detail.Age, Email: passenger_detail.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to create user",
		})
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
