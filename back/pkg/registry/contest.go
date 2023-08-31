package registry

import (
	"back/pkg/adapter/controller"
	"back/pkg/adapter/repository"
	"back/pkg/usecase"
)

func (r *registry) NewContestController() controller.ContestController {
	cr := repository.NewContestRepository(r.db)
	cc := usecase.NewContestUsecase(cr)
	return controller.NewContestController(cc)
}
