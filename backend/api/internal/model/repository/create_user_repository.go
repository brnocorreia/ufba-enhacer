package repository

import (
	"errors"
	"fmt"

	"github.com/brnocorreia/ufba-enhacer/internal/configs/err"
	"github.com/brnocorreia/ufba-enhacer/internal/configs/logger"
	"github.com/brnocorreia/ufba-enhacer/internal/model"
	"github.com/brnocorreia/ufba-enhacer/internal/model/repository/entity/converter"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *err.RestErr) {

	logger.Info("Initializing createUser repository")

	db := ur.databaseConnection
	value := converter.ConvertDomainToEntity(userDomain)

	result := db.Create(&value).Error
	if result != nil {
		logger.Error("Error trying to create user",
			result,
			zap.String("journey", "createUser"),
		)
		if errors.Is(result, gorm.ErrDuplicatedKey) {
			fmt.Println("Está chegando aqui")
			return nil, err.NewBadRequestError("Email already registered. Try another one.")
		}
		fmt.Println("Chegando aqui")
		return nil, err.NewInternalServerError(result.Error(), []err.Causes{})
	}
	// CHECK IF VALUE.ID IS RETURNING CORRECTLY
	logger.Info("createUser repository executed successfully",
		zap.String("userId", value.ID.String()),
		zap.String("journey", "createUser"),
	)
	fmt.Println(value)
	return converter.ConvertEntityToDomain(*value), nil
}
