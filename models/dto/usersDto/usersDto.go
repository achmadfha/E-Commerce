package usersDto

import (
	"github.com/google/uuid"
	"time"
)

type (
	User struct {
		UserID    uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserResponse struct {
		UserID      uuid.UUID   `json:"id"`
		Username    string      `json:"username"`
		Email       string      `json:"email"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		UserProfile UserProfile `json:"profile"`
	}

	UserProfile struct {
		FullName   string `json:"full_name"`
		Address    string `json:"address"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
		Phone      string `json:"phone"`
	}

	UserUpdate struct {
		UserID     uuid.UUID `json:"user_id"`
		FullName   string    `json:"full_name"`
		Address    string    `json:"address"`
		City       string    `json:"city"`
		State      string    `json:"state"`
		Country    string    `json:"country"`
		PostalCode string    `json:"postal_code"`
		Phone      string    `json:"phone"`
	}
)
