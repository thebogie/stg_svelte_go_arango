package repository

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"

	"back/graph/model"
)

type TodoRepository interface {
	Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	Delete(ctx context.Context, id string) (*model.Todo, error)
}

type todorepository struct {
	db driver.Database
}

func NewTodoRepository(db driver.Database) TodoRepository {
	return &todorepository{
		db: db,
	}
}

func (tr *todorepository) Get(ctx context.Context, id string) (*model.Todo, error) {
	return &model.Todo{}, nil
}

func (tr *todorepository) List(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("This is from todo repository")
	return []*model.Todo{}, nil
}

func (tr *todorepository) Create(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{}, nil
}

func (tr *todorepository) Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error) {
	return &model.Todo{}, nil
}

func (tr *todorepository) Delete(ctx context.Context, id string) (*model.Todo, error) {
	return &model.Todo{}, nil
}
