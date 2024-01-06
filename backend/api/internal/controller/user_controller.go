package controller

import (
	"github.com/brnocorreia/ufba-enhacer/internal/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
