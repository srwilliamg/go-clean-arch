package ucuser

type UserInteractor interface {
	GetByName(name string) ([]uint, error)
}

type userInteractor struct {
	UserRepository UserRepository
	UserPresenter  UserPresenter
}

func (us *userInteractor) GetByName(name string) ([]uint, error) {
	u, err := us.UserRepository.FindUserByName(name)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseFindUserByName(u, name), nil
}

func NewUserInteractor(r UserRepository, p UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}
