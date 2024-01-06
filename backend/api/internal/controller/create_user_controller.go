package controller

import (
	"net/http"

	"github.com/brnocorreia/ufba-enhacer/internal/configs/logger"
	"github.com/brnocorreia/ufba-enhacer/internal/configs/validation"
	"github.com/brnocorreia/ufba-enhacer/internal/controller/model/request"
	"github.com/brnocorreia/ufba-enhacer/internal/model"
	"github.com/brnocorreia/ufba-enhacer/internal/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("Initializing CreateUser controller",
		zap.String("journey", "createUser"))

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		UserRequest.Name,
		UserRequest.Email,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call  createUser services",
			err,
			zap.String("journey", "createUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
