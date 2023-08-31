package controller

import (
	"back/graph/model"
	"back/pkg/usecase"
	"context"
)

type GameController interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Game, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
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
