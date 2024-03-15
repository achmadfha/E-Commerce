package authenticationDto

import (
	"github.com/google/uuid"
	"time"
)

type (
	Register struct {
		UsersID   uuid.UUID `json:"users_id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	RegistrationRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	RegistrationResponse struct {
		UsersID   uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
