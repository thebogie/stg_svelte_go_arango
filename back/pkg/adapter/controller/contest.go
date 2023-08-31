package controller

import (
	"context"
	"fmt"

	"back/graph/model"
	"back/pkg/usecase"
)

type ContestController interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Contest, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
}

type contestController struct {
	contestUsecase usecase.ContestUsecase
}

// NewTodoController generates test user controller
func NewContestController(tu usecase.ContestUsecase) ContestController {
	return &contestController{
		contestUsecase: tu,
	}
}

func (tc contestController) List(ctx context.Context) ([]*model.Contest, error) {
	fmt.Println("This is from contest controller")
	return tc.contestUsecase.List(ctx)
}
