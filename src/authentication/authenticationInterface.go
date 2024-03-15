package authentication

import "E-Commerce/models/dto/authenticationDto"

type AuthenticationRepository interface {
	RegistersUsers(req authenticationDto.Register) error
	CheckEmailExists(usrEmail string) (bool, error)
	CheckUsrNameExists(usrName string) (bool, error)
}

type AuthenticationUseCase interface {
	RegisterUsers(req authenticationDto.RegistrationRequest) (authenticationDto.RegistrationResponse, error)
}
