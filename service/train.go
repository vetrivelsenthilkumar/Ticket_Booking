package service

import (
	"Ticket_Booking_App/graph/model"
	"Ticket_Booking_App/initializers"
	"context"
)

func TrainCreate(ctx context.Context, input model.NewTrain) (*model.Train, error) {
	db := initializers.DB

	train := model.Train{
		TrainNumber: input.TrainNumber,
		From:        input.From,
		To:          input.To,
		CoachNumber: input.CoachNumber,
		SeatNumber:  input.SeatNumber,
		SeatType:    input.SeatType,
	}

	if err := db.Model(train).Create(&train).Error; err != nil {
		return nil, err
	}

	return &train, nil
}

func TrainGetByNUmber(ctx context.Context, train_number string) (*model.Train, error) {
	db := initializers.DB

	var train model.Train
	if err := db.Model(train).Where("train_number LIKE ?", train_number).Take(&train).Error; err != nil {
		return nil, err
	}

	return &train, nil
}
