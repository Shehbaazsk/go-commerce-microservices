package models

import (
	"time"

	"github.com/shehbaazsk/go-commerce-microservices/pkg/response"
)

type CreateRoleRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

type UpdateRoleRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

type RoleResponse struct {
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	IsActive    *bool     `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListRoleRequest struct {
	response.PaginationInput
}
