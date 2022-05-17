package appcontroller

import useradapter "go-clean-arch/interface/user"

type AppController struct {
	User interface {
		useradapter.UserInterfaceController
	}
}

func NewAppController() *AppController {
	return &AppController{}
}

func (ap *AppController) AddUserController(uc *useradapter.UserInterfaceController) {
	ap.User = *uc
}
