package model

type UserDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string

	SetID(string)
}

func NewUserDomain(name, email string) UserDomainInterface {
	return &userDomain{
		name:  name,
		email: email,
	}
}

func NewUserLoginDomain(id, email string) UserDomainInterface {
	return &userDomain{
		id:    id,
		email: email,
	}
}

type userDomain struct {
	id    string
	name  string
	email string
}

func (ud *userDomain) SetID(id string) {
	ud.id = id

}

func (ud *userDomain) GetID() string {
	return ud.id

}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
