package service

import (
	"github.com/brnocorreia/ufba-enhacer/internal/configs/err"
	"github.com/brnocorreia/ufba-enhacer/internal/configs/logger"
	"github.com/brnocorreia/ufba-enhacer/internal/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *err.RestErr) {

	logger.Info("Initializing CreateUser model",
		zap.String("journey", "createUser"))

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
