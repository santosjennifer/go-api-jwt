package service

import (
	"go-api-jwt/src/configuration/rest_error"
	"go-api-jwt/src/model"
	"go-api-jwt/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_error.RestErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_error.RestErr)
	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_error.RestErr)

	UpdateUser(string, model.UserDomainInterface) *rest_error.RestErr
	DeleteUser(string) *rest_error.RestErr

	LoginUserServices(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, *rest_error.RestErr)
}
