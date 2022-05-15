package user

import (
	reqctx "go-clean-arch/interface/request-context"
	usecases "go-clean-arch/use-cases/user"
	"net/http"
)

type userInterfaceController struct {
	userInteractor usecases.UserInteractor
}

type UserInterfaceController interface {
	GetByName(c reqctx.Context) error
}

func (uc *userInterfaceController) GetByName(c reqctx.Context) error {
	name := c.QueryParam("name")

	u, err := uc.userInteractor.GetByName(name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func NewUseInterfaceController(us usecases.UserInteractor) UserInterfaceController {
	return &userInterfaceController{userInteractor: us}
}
