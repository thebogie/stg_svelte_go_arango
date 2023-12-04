package usecase

import (
	"back/graph/model"
	"back/pkg/adapter/repository"
	"context"
	"fmt"
)

type contestUsecase struct {
	contestRepository repository.ContestRepository
}

type ContestUsecase interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Contest, error)
	GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
}

func NewContestUsecase(cr repository.ContestRepository) ContestUsecase {
	return &contestUsecase{
		contestRepository: cr,
	}
}

func (cu contestUsecase) List(ctx context.Context) ([]*model.Contest, error) {

	return nil, nil
	//tu.todoRepository.List(ctx)
}

func (cu contestUsecase) GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error) {
	fmt.Println("This is from contest usecase")

	return cu.contestRepository.GetContestsPlayerTotalResults(ctx, player)
}
