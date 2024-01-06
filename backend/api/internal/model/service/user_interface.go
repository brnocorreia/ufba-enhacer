package service

import (
	"github.com/brnocorreia/ufba-enhacer/internal/configs/err"
	"github.com/brnocorreia/ufba-enhacer/internal/model"
	"github.com/brnocorreia/ufba-enhacer/internal/model/repository"
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
		model.UserDomainInterface,
		*err.RestErr,
	)
}
