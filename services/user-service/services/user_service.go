package services

import (
	"context"

	"github.com/shehbaazsk/go-commerce-microservices/pkg/converters"
	db "github.com/shehbaazsk/go-commerce-microservices/user-service/db/queries"
	"github.com/shehbaazsk/go-commerce-microservices/user-service/models"
)

func mapUserToResponse(user db.User) models.UserResponse {
	return models.UserResponse{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  converters.TextOrNil(user.LastName),
		Email:     user.Email,
		PhoneNo:   converters.TextOrNil(user.PhoneNo),
		DOB:       converters.DateOrNil(user.Dob),
		IsActive:  converters.BoolOrNil(user.IsActive),
	}
}

type UserService interface {
	CreateUser(ctx context.Context, tx *db.Queries, req db.CreateUserParams) (models.UserResponse, error)
	UpdateUser(ctx context.Context, tx *db.Queries, req db.UpdateUserParams) (models.UserResponse, error)
	DeleteUser(ctx context.Context, tx *db.Queries, id int) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(ctx context.Context, tx *db.Queries, req db.CreateUserParams) (models.UserResponse, error) {
	user, err := tx.CreateUser(ctx, req)
	if err != nil {
		return models.UserResponse{}, err
	}
	return mapUserToResponse(user), nil
}

func (s *userService) UpdateUser(ctx context.Context, q *db.Queries, req db.UpdateUserParams) (models.UserResponse, error) {
	user, err := q.UpdateUser(ctx, req)
	if err != nil {
		return models.UserResponse{}, err
	}
	return mapUserToResponse(user), nil
}

func (s *userService) DeleteUser(ctx context.Context, q *db.Queries, id int) error {
	return q.DeleteUser(ctx, int32(id))
}
