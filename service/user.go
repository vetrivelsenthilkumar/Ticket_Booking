package service

import (
	"Ticket_Booking_App/graph/model"
	"Ticket_Booking_App/initializers"
	"Ticket_Booking_App/tools"
	"context"
	"strings"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := initializers.DB

	*input.Password = tools.HashPassword(*input.Password)

	user := model.User{
		Name:     input.Name,
		Age:      input.Age,
		Email:    strings.ToLower(input.Email),
		Password: *input.Password,
	}

	if err := db.Model(user).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// func UserGetByID(ctx context.Context, id string) (*model.User, error) {
// 	db := initializers.DB

// 	var user model.User
// 	if err := db.Model(user).Where("id = ?", id).Take(&user).Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	db := initializers.DB

	var user model.User
	if err := db.Model(user).Where("email LIKE ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
