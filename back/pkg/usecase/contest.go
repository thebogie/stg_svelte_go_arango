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
	List(ctx context.Context) ([]*model.Contest, error)
	GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error)
	GetStats(ctx context.Context, player string) (*model.Stats, error)
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

func (cu contestUsecase) GetStats(ctx context.Context, player string) (*model.Stats, error) {
	fmt.Println("Contest Usecase: GetStats")
	//response, error := cu.contestRepository.GetContestsPlayerTotalResults(ctx, player)

	return nil, nil
}
