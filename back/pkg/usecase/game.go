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
	List(ctx context.Context) ([]*model.Game, error)
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
