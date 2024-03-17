package usersUseCase

import (
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/usersDto"
	"E-Commerce/src/users"
	"database/sql"
	"errors"
	"math"
)

type userUC struct {
	userRepository users.UserRepository
}

func NewUserUseCase(userRepo users.UserRepository) users.UserUseCase {
	return &userUC{userRepo}
}

func (u userUC) RetrieveAllUsers(page, pageSize int) ([]usersDto.User, json.Pagination, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 5
	}

	userData, err := u.userRepository.RetrieveAllUsers(page, pageSize)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, json.Pagination{}, errors.New("no rows found")
		}
		return nil, json.Pagination{}, err

	}

	totalUserData, err := u.userRepository.CountAllUsers()
	if err != nil {
		return nil, json.Pagination{}, err
	}

	totalPages := int(math.Ceil(float64(totalUserData) / float64(pageSize)))
	if page > totalPages {
		return nil, json.Pagination{}, errors.New("01")
	}

	if totalPages == 0 && totalUserData > 0 {
		totalPages = 1
	}

	pagination := json.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: totalUserData,
	}

	return userData, pagination, nil
}

func (u userUC) RetrieveUsersByID(usrID string) (usersDto.UserResponse, error) {
	profileExists, err := u.userRepository.CheckUserProfileExists(usrID)
	if err != nil {
		return usersDto.UserResponse{}, err
	}

	if !profileExists {
		// profile doesn't exists
		return usersDto.UserResponse{}, errors.New("01")
	}

	userData, err := u.userRepository.RetrieveUsersByID(usrID)
	if err != nil {
		return usersDto.UserResponse{}, err
	}

	return userData, nil
}
