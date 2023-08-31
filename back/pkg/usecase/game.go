package usecase

import (
	"back/graph/model"
	"back/pkg/adapter/repository"
	"context"
)

type gameUsecase struct {
	gameRepository repository.GameRepository
}

type GameUsecase interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Game, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
}

func NewGameUsecase(ur repository.GameRepository) GameUsecase {
	return &gameUsecase{
		gameRepository: ur,
	}
}

func (gu gameUsecase) List(ctx context.Context) ([]*model.Game, error) {

	return gu.gameRepository.List(ctx)
	//tu.todoRepository.List(ctx)
}
