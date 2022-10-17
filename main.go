package main

import (
	"Ticket_Booking_App/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connecting_db()
}

// func Main() {
// 	r := gin.Default()
// 	r.SetTrustedProxies([]string{"192.168.1.2"})
// 	r.POST("/signup", controllers.Signup)
// 	r.POST("/login", controllers.Login)
// 	r.GET("/users", controllers.GetUsers)
// 	r.PUT("/users/:id", controllers.UsersUpdate)
// 	r.DELETE("/users/:id", controllers.UsersDelete)
// 	r.GET("/validate", controllers.Validate)
// 	r.POST("/train", controllers.BookTrain)
// 	r.GET("/train", controllers.TrainDetails)
// 	r.PUT("/train/:id", controllers.TrainUpdate)
// 	r.DELETE("/train/:id", controllers.CancelBooking)
// 	r.Run()
// }
