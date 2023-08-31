package controller

import (
	"context"
	"fmt"

	"back/graph/model"
	"back/pkg/usecase"
)

type TodoController interface {
	Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	Delete(ctx context.Context, id string) (*model.Todo, error)
}

type todoController struct {
	todoUsecase usecase.TodoUsecase
}

// NewTodoController generates test user controller
func NewTodoController(tu usecase.TodoUsecase) TodoController {
	return &todoController{
		todoUsecase: tu,
	}
}

func (tc todoController) Get(ctx context.Context, id string) (*model.Todo, error) {
	return tc.todoUsecase.Get(ctx, id)
}

func (tc todoController) List(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("This is from todo controller")
	return tc.todoUsecase.List(ctx)
}

func (tc todoController) Create(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return tc.todoUsecase.Create(ctx, input)
}

func (tc todoController) Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error) {
	return tc.todoUsecase.Update(ctx, input)
}

func (tc todoController) Delete(ctx context.Context, id string) (*model.Todo, error) {
	return tc.todoUsecase.Delete(ctx, id)
}
