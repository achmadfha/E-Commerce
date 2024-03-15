package validation

import (
	"E-Commerce/models/dto/authenticationDto"
	"E-Commerce/models/dto/json"
	"strings"
	"unicode"
)

func ValidateRegister(req authenticationDto.RegistrationRequest) []json.ValidationField {
	var validationErrors []json.ValidationField

	if req.Username == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "username",
			Message:   "Username cannot be empty",
		})
	}

	if req.Password == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "password",
			Message:   "Password cannot be empty",
		})
	} else {
		if len(req.Password) < 8 {
			validationErrors = append(validationErrors, json.ValidationField{
				FieldName: "password",
				Message:   "Password must be at least 8 characters long",
			})
		}

		hasUppercase := false
		hasLowercase := false
		hasDigit := false
		for _, char := range req.Password {
			switch {
			case unicode.IsUpper(char):
				hasUppercase = true
			case unicode.IsLower(char):
				hasLowercase = true
			case unicode.IsDigit(char):
				hasDigit = true
			}
		}
		if !hasUppercase || !hasLowercase || !hasDigit {
			validationErrors = append(validationErrors, json.ValidationField{
				FieldName: "password",
				Message:   "Password must contain at least one uppercase letter, one lowercase letter, and one digit",
			})
		}
	}

	if req.Email == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "email",
			Message:   "Email cannot be empty",
		})
	} else {
		parts := strings.Split(req.Email, "@")
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			validationErrors = append(validationErrors, json.ValidationField{
				FieldName: "email",
				Message:   "Invalid email format",
			})
		}
	}

	return validationErrors
}
