package authenticationUseCase

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/authenticationDto"
	"E-Commerce/pkg/utils"
	"E-Commerce/src/authentication"
	"errors"
	"github.com/google/uuid"
	"time"
)

type authenticationUC struct {
	authenticationRepository authentication.AuthenticationRepository
}

func NewAuthenticationUseCase(authentication authentication.AuthenticationRepository) authentication.AuthenticationUseCase {
	return &authenticationUC{authentication}
}

func (a authenticationUC) RegisterUsers(req authenticationDto.RegistrationRequest) (authenticationDto.RegistrationResponse, error) {

	emailExists, err := a.authenticationRepository.CheckEmailExists(req.Email)
	if err != nil {
		return authenticationDto.RegistrationResponse{}, err
	}
	if emailExists {
		return authenticationDto.RegistrationResponse{}, errors.New("01")
	}

	usernameExists, err := a.authenticationRepository.CheckUsrNameExists(req.Username)
	if err != nil {
		return authenticationDto.RegistrationResponse{}, err
	}
	if usernameExists {
		return authenticationDto.RegistrationResponse{}, errors.New("02")
	}

	usrID, err := uuid.NewRandom()
	if err != nil {
		return authenticationDto.RegistrationResponse{}, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	usrRole := constants.DefaultRole
	currentTime := time.Now()

	usrData := authenticationDto.Register{
		UsersID:   usrID,
		Username:  req.Username,
		Password:  hashedPassword,
		Email:     req.Email,
		Role:      usrRole,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	err = a.authenticationRepository.RegistersUsers(usrData)
	if err != nil {
		return authenticationDto.RegistrationResponse{}, err
	}

	usrRes := authenticationDto.RegistrationResponse{
		UsersID:   usrID,
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: currentTime,
	}

	return usrRes, nil
}
