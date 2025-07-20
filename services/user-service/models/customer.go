package models

import "time"

type CreateCustomerRequest struct {
	FirstName string     `json:"first_name" binding:"required"`
	LastName  *string    `json:"last_name"`
	Email     string     `json:"email" binding:"required,email"`
	Password  string     `json:"password" binding:"required"`
	PhoneNo   *string    `json:"phone_no,omitempty"`
	DOB       *time.Time `json:"dob,omitempty"`
}

type UpdateCustomerRequest struct {
	FirstName *string    `json:"first_name"`
	LastName  *string    `json:"last_name"`
	Email     *string    `json:"email"`
	PhoneNo   *string    `json:"phone_no"`
	DOB       *time.Time `json:"dob"`
	IsActive  *bool      `json:"is_active"`
}

type CustomerResponse struct {
	UserID    int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  *string    `json:"last_name"`
	Email     string     `json:"email"`
	PhoneNo   *string    `json:"phone_no"`
	DOB       *time.Time `json:"dob"`
}

type ListCustomerRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func (r *ListCustomerRequest) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PerPage <= 0 {
		r.PerPage = 10
	}
}
