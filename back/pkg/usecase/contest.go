package usecase

import (
	"back/graph/model"
	"back/pkg/adapter/repository"
	"context"
)

type contestUsecase struct {
	contestRepository repository.ContestRepository
}

type ContestUsecase interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Contest, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
}

func NewContestUsecase(tr repository.ContestRepository) ContestUsecase {
	return &contestUsecase{
		contestRepository: tr,
	}
}

func (tu contestUsecase) List(ctx context.Context) ([]*model.Contest, error) {

	return nil, nil
	//tu.todoRepository.List(ctx)
}
