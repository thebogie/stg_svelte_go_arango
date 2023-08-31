package repository

import (
	"back/graph/model"
	"context"
	"github.com/arangodb/go-driver"
)

type ContestRepository interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Contest, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
}

type contestrepository struct {
	db driver.Database
}

func NewContestRepository(db driver.Database) ContestRepository {
	return &contestrepository{
		db: db,
	}
}

func (tr *contestrepository) List(ctx context.Context) ([]*model.Contest, error) {
	//var retuser = &model.UserData{}

	return []*model.Contest{}, nil
}
