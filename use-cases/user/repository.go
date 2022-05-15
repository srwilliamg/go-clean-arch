package ucuser

import "go-clean-arch/entities"

type UserRepository interface {
	FindUserByName(name string) ([]*entities.User, error)
}
