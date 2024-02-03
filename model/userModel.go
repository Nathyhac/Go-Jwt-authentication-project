package model

import "time"

type User struct {
	ID           *string    `json:"id"`
	FirstName    *string    `json:"first_name"`
	LastName     *string    `json:"last_name"`
	Password     *string    `json:"password"`
	Email        *string    `json:"email"`
	Phone        *string    `json:"phone"`
	Token        *string    `json:"token"`
	UserType     *string    `json:"user_type"`
	RefreshToken *string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       string    `json:"user_id"`
}