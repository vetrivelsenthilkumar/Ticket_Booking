package main

import (
	// "postgres/controllers"

	"postgres/controllers"
	"postgres/initializers"
	"postgres/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connecting_db()
	initializers.SyncDatabase()
}

func main() {
	println("Hello")
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	// r.POST("/users", controllers.UsersCreate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUsersById)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.DELETE("/users/:id", controllers.UsersDelete)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/train", controllers.BookTrain)
	r.GET("/train", controllers.TrainDetails)
	r.PUT("/train/:train_number", controllers.TrainUpdate)
	r.DELETE("/train/:train_number", controllers.CancelBooking)
	r.Run()
}
