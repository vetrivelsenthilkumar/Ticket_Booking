package service

import (
	"Ticket_Booking_App/graph/model"
	"Ticket_Booking_App/initializers"
	"context"
	"fmt"
	"net/smtp"
	"os"

	"gorm.io/gorm"
)

func BookTrain(ctx context.Context, input model.NewTrain) (interface{}, error) {
	db := initializers.DB
	_, err := TrainGetByNUmber(ctx, input.TrainNumber)

	train := model.Train{
		TrainNumber: input.TrainNumber,
		From:        input.From,
		To:          input.To,
		CoachNumber: input.CoachNumber,
		SeatNumber:  input.SeatNumber,
		SeatType:    input.SeatType,
	}
	if err == nil {
		// if err != record not found
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	if err := db.Model(train).Create(&train).Error; err != nil {
		return nil, err
	}

	Mail("vetrisenthilmkce@gmail.com", "Ticket is Booked Successfully.")

	return &train, nil

}

func Mail(Email, Msg string) {

	from := "vetrisenthilmkce@gmail.com"
	password := os.Getenv("PASSWD")
	ToEmail := Email
	toList := []string{ToEmail}

	host := "smtp.gmail.com"

	port := "587"

	msg := Msg

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
