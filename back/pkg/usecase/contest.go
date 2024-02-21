package usecase

import (
	"back/graph/model"
	"back/pkg/adapter/repository"
	"back/pkg/utils"
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
	CreateContest(ctx context.Context, newContest model.InputContest) (string, error)
}

func NewContestUsecase(cr repository.ContestRepository) ContestUsecase {
	return &contestUsecase{
		contestRepository: cr,
	}
}

func (cu contestUsecase) CreateContest(ctx context.Context, newContest model.InputContest) (string, error) {
	utils.PrintFunctionName()

	contestId, _ := cu.contestRepository.CreateContest(ctx, newContest)

	return contestId, nil
	//tu.todoRepository.List(ctx)
}

func (cu contestUsecase) List(ctx context.Context) ([]*model.Contest, error) {

	return nil, nil
	//tu.todoRepository.List(ctx)
}

func (cu contestUsecase) GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error) {
	utils.PrintFunctionName()

	return cu.contestRepository.GetContestsPlayerTotalResults(ctx, player)
}

func (cu contestUsecase) GetStats(ctx context.Context, player string) (*model.Stats, error) {
	fmt.Println("Contest Usecase: GetStats")
	//response, error := cu.contestRepository.GetContestsPlayerTotalResults(ctx, player)

	return nil, nil
}
