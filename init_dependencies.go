package main

import (
	"go-api-jwt/src/controller"
	"go-api-jwt/src/model/repository"
	"go-api-jwt/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
