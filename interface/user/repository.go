package user

import (
	"go-clean-arch/data"
	"go-clean-arch/entities"
	usecases "go-clean-arch/use-cases/user"
)

type interfaceRepositoryUser struct {
	db *data.DataObject
}

func NewUserInterfaceRepository(db *data.DataObject) usecases.UserRepository {
	return &interfaceRepositoryUser{db}
}

func (iru *interfaceRepositoryUser) FindUserByName(name string) ([]*entities.User, error) {
	data := iru.db.FindUserByName(name)

	return data, nil
}
