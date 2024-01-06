package converter

import (
	"github.com/brnocorreia/ufba-enhacer/internal/model"
	"github.com/brnocorreia/ufba-enhacer/internal/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.User,
) model.UserDomainInterface {

	domain := model.NewUserDomain(entity.Name, entity.Email)

	domain.SetID(entity.ID.String())
	return domain
}
