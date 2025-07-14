package models

import "time"

type CreateUserReq struct {
	FirstName string     `json:"first_name" binding:"required"`
	LastName  *string    `json:"last_name,omitempty"`
	Email     string     `json:"email" binding:"required,email"`
	Password  string     `json:"password" binding:"required"`
	PhoneNo   *string    `json:"phone_no,omitempty"`
	DOB       *time.Time `json:"dob,omitempty"`
}

type UserResponse struct {
	ID        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  *string    `json:"last_name"`
	Email     string     `json:"email"`
	PhoneNo   *string    `json:"phone_no"`
	DOB       *time.Time `json:"dob"`
	IsActive  *bool      `json:"is_active"`
}

type UpdateUserRequest struct {
	FirstName *string    `json:"first_name"`
	Email     *string    `json:"email" binding:"omitempty,email"`
	LastName  *string    `json:"last_name"`
	PhoneNo   *string    `json:"phone_no"`
	DOB       *time.Time `json:"dob"`
	IsActive  *bool      `json:"is_active"`
}
