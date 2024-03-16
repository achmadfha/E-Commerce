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
)
