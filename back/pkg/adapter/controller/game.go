package controller

import (
	"back/graph/model"
	"back/pkg/usecase"
	"context"
)

type GameController interface {
	List(ctx context.Context) ([]*model.Game, error)
}

type gameController struct {
	gameUsecase usecase.GameUsecase
}

// NewTodoController generates test user controller
func NewGameController(gu usecase.GameUsecase) GameController {
	return &gameController{
		gameUsecase: gu,
	}
}

func (gc gameController) List(ctx context.Context) ([]*model.Game, error) {

	return gc.gameUsecase.List(ctx)
}
