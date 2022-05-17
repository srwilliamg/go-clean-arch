package user

import (
	"go-clean-arch/entities"
	database "go-clean-arch/infrastructure/db"
	usecases "go-clean-arch/use-cases/user"
)

type interfaceRepositoryUser struct {
	db *database.DataObject
}

func NewUserInterfaceRepository(db *database.DataObject) usecases.UserRepository {
	return &interfaceRepositoryUser{db}
}

func (iru *interfaceRepositoryUser) FindUserByName(name string) ([]*entities.User, error) {
	data := iru.db.MultipleUserByName(name)

	return data, nil
}
