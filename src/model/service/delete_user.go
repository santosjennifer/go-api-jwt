package service

import (
	"go-api-jwt/src/configuration/logger"
	"go-api-jwt/src/configuration/rest_error"

	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string) *rest_error.RestErr {

	logger.Info("Init deleteUser model.",
		zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info(
		"deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))
	return nil
}
