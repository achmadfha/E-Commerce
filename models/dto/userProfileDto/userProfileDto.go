package userProfileDto

import (
	"github.com/google/uuid"
	"time"
)

type (
	User struct {
		UsersID     uuid.UUID   `json:"users_id"`
		Username    string      `json:"username"`
		Email       string      `json:"email"`
		UserProfile UserProfile `json:"userProfile"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
	}

	UserProfile struct {
		UserProfileID uuid.UUID `json:"user_profile_id"`
		FullName      string    `json:"full_name"`
		Address       string    `json:"address"`
		City          string    `json:"city"`
		State         string    `json:"state"`
		Country       string    `json:"country"`
		PostalCode    string    `json:"postalCode"`
		Phone         string    `json:"phone"`
	}
)
