package user

import (
	"go-clean-arch/entities"
	usecases "go-clean-arch/use-cases/user"
)

type userInterfacePresenter struct{}

func (up *userInterfacePresenter) ResponseFindUserByName(users []*entities.User, name string) []uint {
	var response = []uint{}

	for _, v := range users {
		if v.Name == name {
			response = append(response, v.ID)
		}
	}

	return response
}

func NewUserPresenter() usecases.UserPresenter {
	return &userInterfacePresenter{}
}
