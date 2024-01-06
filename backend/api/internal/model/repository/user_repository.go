package repository

import (
	"github.com/brnocorreia/ufba-enhacer/internal/configs/err"
	"github.com/brnocorreia/ufba-enhacer/internal/model"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *gorm.DB
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *err.RestErr)
}
