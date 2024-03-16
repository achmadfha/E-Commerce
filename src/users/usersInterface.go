package users

import (
	"E-Commerce/models/dto/json"
	"E-Commerce/models/dto/usersDto"
)

type UserRepository interface {
	RetrieveAllUsers(page, pageSize int) ([]usersDto.User, error)
	CountAllUsers() (int, error)
}

type UserUseCase interface {
	RetrieveAllUsers(page, pageSize int) ([]usersDto.User, json.Pagination, error)
}
