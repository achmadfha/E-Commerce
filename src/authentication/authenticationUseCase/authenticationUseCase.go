package authenticationUseCase

import (
	"E-Commerce/models/constants"
	"E-Commerce/models/dto/authenticationDto"
	"E-Commerce/pkg/utils"
	"E-Commerce/src/authentication"
	"errors"
	"fmt"
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

func (a authenticationUC) LoginUsers(req authenticationDto.LoginRequest) (token string, err error) {
	usr, err := a.authenticationRepository.RetrieveUsers(req.Email)
	if err != nil {
		// 01 email not registered
		if err.Error() == "01" {
			return "", errors.New("01")
		}
		return "", err
	}

	if err := utils.VerifyPassword(usr.Password, req.Password); err != nil {
		// 02 password didn't match
		fmt.Println(err)
		return "", errors.New("02")
	}

	token, err = utils.GenerateToken(usr.UsersID)
	if err != nil {
		return "", err
	}

	return token, nil
}
