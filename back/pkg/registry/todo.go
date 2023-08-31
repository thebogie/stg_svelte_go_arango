package registry

import (
	"back/pkg/adapter/controller"
	"back/pkg/adapter/repository"
	"back/pkg/usecase"
)

func (r *registry) NewTodoController() controller.TodoController {
	tr := repository.NewTodoRepository(r.db)
	uc := usecase.NewTodoUsecase(tr)
	return controller.NewTodoController(uc)
}
