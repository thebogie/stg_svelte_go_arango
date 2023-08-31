package registry

import (
	"back/pkg/adapter/controller"
	"back/pkg/adapter/repository"
	"back/pkg/usecase"
)

func (r *registry) NewUserController() controller.UserController {
	ur := repository.NewUserRepository(r.db)
	uc := usecase.NewUserUsecase(ur)
	return controller.NewUserController(uc)
}
