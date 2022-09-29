package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"postgres/initializers"
	"postgres/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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
	mail(passenger_detail.Email, "Successfully created the account")
}

func Login(c *gin.Context) {

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
	var user models.Passenger
	initializers.DB.First(&user, "email = ?", passenger_detail.Email)

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passenger_detail.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to Compare password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to Create token",
		})
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.Passenger
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUsersById(c *gin.Context) {

	id := c.Param("id")

	var user []models.Passenger
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersUpdate(c *gin.Context) {

	id := c.Param("id")

	var passenger_detail struct {
		Id       uint
		Name     string
		Age      int
		Email    string
		Password string
	}

	c.Bind(&passenger_detail)

	var user []models.Passenger
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.Passenger{Name: passenger_detail.Name, Age: passenger_detail.Age, Email: passenger_detail.Email, Password: passenger_detail.Password})

	c.JSON(200, gin.H{
		"user": user,
	})
	mail(passenger_detail.Email, "The passsenger details are updated.")
}

func UsersDelete(c *gin.Context) {

	id := c.Param("id")

	var user []models.Passenger
	initializers.DB.Delete(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
	c.Status(200)

}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func BookTrain(c *gin.Context) {
	var train struct {
		Train_number uint
		From         string
		To           string
		Coach_number string
		Seat_number  int
		Seat_type    string
	}

	c.Bind(&train)

	trains := models.Train{Train_number: uint(train.Train_number), From: train.From, To: train.To, Coach_number: train.Coach_number, Seat_number: train.Seat_number, Seat_type: train.Seat_type}

	result := initializers.DB.Create(&trains)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"trains": trains,
	})

	mail("vetrisenthilmkce@gmail.com", "Ticket is Booked Successfully.")
}

func TrainDetails(c *gin.Context) {
	var trains []models.Train
	initializers.DB.Find(&trains)

	c.JSON(200, gin.H{
		"trains": trains,
	})
}

func TrainUpdate(c *gin.Context) {

	train_number := c.Param("train_number")

	var train struct {
		Train_number uint
		From         string
		To           string
		Coach_number string
		Seat_number  int
		Seat_type    string
	}

	c.Bind(&train)

	var trains []models.Train
	initializers.DB.First(&trains, train_number)

	initializers.DB.Model(&trains).Updates(models.Train{Train_number: uint(train.Train_number), From: train.From, To: train.To, Coach_number: train.Coach_number, Seat_number: train.Seat_number, Seat_type: train.Seat_type})

	c.JSON(200, gin.H{
		"trains": trains,
	})

	mail("vetrisenthilmkce@gmail.com", "Train details are updated Successfully.")
}

func CancelBooking(c *gin.Context) {

	train_number := c.Param("train_number")

	var user []models.Passenger
	initializers.DB.Delete(&user, train_number)

	c.JSON(200, gin.H{
		"user": user,
	})
	c.Status(200)

	mail("vetrisenthilmkce@gmail.com", "Ticket is cancelled Successfully.")
}

func mail(Email, Msg string) {
	from := os.Getenv("MAIL")
	password := os.Getenv("PASSWD")
	toEmail := Email
	to := []string{toEmail}

	host := "smtp.gmail.com"
	port := "587"
	msg := Msg
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, to, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to the passenger")
}
