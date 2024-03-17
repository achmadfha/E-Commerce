package users

import (
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/usersDto"
)

type UserRepository interface {
	RetrieveAllUsers(page, pageSize int) ([]usersDto.User, error)
	CountAllUsers() (int, error)
	CheckUserProfileExists(usrID string) (bool, error)
	RetrieveUsersByID(usrID string) (usrData usersDto.UserResponse, err error)
}

type UserUseCase interface {
	RetrieveAllUsers(page, pageSize int) ([]usersDto.User, json.Pagination, error)
	RetrieveUsersByID(usrID string) (usersDto.UserResponse, error)
}
