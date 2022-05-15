package ucuser

import "go-clean-arch/entities"

type UserPresenter interface {
	ResponseFindUserByName([]*entities.User, string) []uint
}
