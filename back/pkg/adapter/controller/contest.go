package controller

import (
	"context"
	"fmt"

	"back/graph/model"
	"back/pkg/usecase"
)

type ContestController interface {
	List(ctx context.Context) ([]*model.Contest, error)
	GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error)
	GetStats(ctx context.Context, player string) (*model.Stats, error)
	CreateContest(ctx context.Context, newContest model.InputContest) (string, error)
}

type contestController struct {
	contestUsecase usecase.ContestUsecase
}

// NewTodoController generates test user controller
func NewContestController(cu usecase.ContestUsecase) ContestController {
	return &contestController{
		contestUsecase: cu,
	}
}
func (tc contestController) CreateContest(ctx context.Context, newContest model.InputContest) (string, error) {
	fmt.Println("This is from CreateContest controller")
	contestId, err := tc.contestUsecase.CreateContest(ctx, newContest)
	if err != nil {
		return "", err
	}

	return contestId, nil
}

func (tc contestController) List(ctx context.Context) ([]*model.Contest, error) {
	fmt.Println("This is from contest controller")
	return tc.contestUsecase.List(ctx)
}

func (tc contestController) GetStats(ctx context.Context, player string) (*model.Stats, error) {
	fmt.Println("Contest Controller: GetStats")
	return tc.contestUsecase.GetStats(ctx, player)
}

func (tc contestController) GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error) {
	fmt.Println("This is from contest controller")
	var test, err = tc.contestUsecase.GetContestsPlayerTotalResults(ctx, player)

	fmt.Println("This is from contest controller return %+v\n", test)
	return test, err
}
