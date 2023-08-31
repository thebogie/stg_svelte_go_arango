package registry

import (
	"back/pkg/adapter/controller"
	"back/pkg/adapter/repository"
	"back/pkg/usecase"
)

func (r *registry) NewGameController() controller.GameController {
	gr := repository.NewGameRepository(r.db)
	gc := usecase.NewGameUsecase(gr)
	return controller.NewGameController(gc)
}
