package view

import (
	"github.com/brnocorreia/ufba-enhacer/internal/controller/model/response"
	"github.com/brnocorreia/ufba-enhacer/internal/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
	}
}
