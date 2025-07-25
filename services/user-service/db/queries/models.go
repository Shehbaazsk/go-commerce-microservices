// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type City struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	StateID int32  `json:"state_id"`
}

type Country struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Permission struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
	IsActive    pgtype.Bool      `json:"is_active"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type RolePermission struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

type State struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	CountryID int32  `json:"country_id"`
}

type User struct {
	ID           int32            `json:"id"`
	FirstName    string           `json:"first_name"`
	LastName     pgtype.Text      `json:"last_name"`
	Email        string           `json:"email"`
	PasswordHash string           `json:"password_hash"`
	PhoneNo      pgtype.Text      `json:"phone_no"`
	Dob          pgtype.Date      `json:"dob"`
	IsActive     pgtype.Bool      `json:"is_active"`
	IsDeleted    pgtype.Bool      `json:"is_deleted"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

type UserAddress struct {
	ID        int32            `json:"id"`
	UserID    int32            `json:"user_id"`
	Label     pgtype.Text      `json:"label"`
	Address1  string           `json:"address1"`
	Address2  pgtype.Text      `json:"address2"`
	City      pgtype.Text      `json:"city"`
	State     pgtype.Text      `json:"state"`
	Country   pgtype.Text      `json:"country"`
	PinCode   pgtype.Text      `json:"pin_code"`
	IsDefault pgtype.Bool      `json:"is_default"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type UserRole struct {
	ID        int32            `json:"id"`
	UserID    int32            `json:"user_id"`
	RoleID    int32            `json:"role_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}
