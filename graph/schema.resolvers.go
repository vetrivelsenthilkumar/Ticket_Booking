package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Ticket_Booking_App/graph/generated"
	"Ticket_Booking_App/graph/model"
	"Ticket_Booking_App/service"
	"context"
	"fmt"
)

// Login is the resolver for the login field.
func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (interface{}, error) {
	return service.UserLogin(ctx, email, password)
}

// Register is the resolver for the register field.
func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.NewUser) (interface{}, error) {
	return service.UserRegister(ctx, input)
}

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

// Train is the resolver for the train field.
func (r *mutationResolver) Train(ctx context.Context) (*model.TrainReserve, error) {
	return &model.TrainReserve{}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, email string) (*model.User, error) {
	return service.UserGetByEmail(ctx, email)
}

// Train is the resolver for the train field.
func (r *queryResolver) Train(ctx context.Context, trainNumber string) (*model.Train, error) {
	return service.TrainGetByNUmber(ctx, trainNumber)
}

// Protected is the resolver for the protected field.
func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Protected - protected"))
}

// Book is the resolver for the book field.
func (r *trainReserveResolver) Book(ctx context.Context, obj *model.TrainReserve, input model.NewTrain) (interface{}, error) {
	return service.BookTrain(ctx, input)
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TrainReserve returns generated.TrainReserveResolver implementation.
func (r *Resolver) TrainReserve() generated.TrainReserveResolver { return &trainReserveResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type trainReserveResolver struct{ *Resolver }
