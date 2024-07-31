package repository

import (
	"go-api-jwt/src/configuration/rest_error"
	"go-api-jwt/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_error.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_error.RestErr

	DeleteUser(
		userId string,
	) *rest_error.RestErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_error.RestErr)
	FindUserByEmailAndPassword(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_error.RestErr)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_error.RestErr)
}
