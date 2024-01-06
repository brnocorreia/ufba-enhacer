package converter

import (
	"github.com/brnocorreia/ufba-enhacer/internal/model"
	"github.com/brnocorreia/ufba-enhacer/internal/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.User {
	return &entity.User{
		Name:  domain.GetName(),
		Email: domain.GetEmail(),
	}
}
