package usecase

import (
	"context"
	"fmt"

	"back/graph/model"
	"back/pkg/adapter/repository"
)

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

type TodoUsecase interface {
	Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	Delete(ctx context.Context, id string) (*model.Todo, error)
}

func NewTodoUsecase(tr repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepository: tr,
	}
}

func (tu todoUsecase) Get(ctx context.Context, id string) (*model.Todo, error) {
	return tu.todoRepository.Get(ctx, id)
}

func (tu todoUsecase) List(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("This is from todo usecase")
	return tu.todoRepository.List(ctx)
}

func (tu todoUsecase) Create(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return tu.todoRepository.Create(ctx, input)
}

func (tu todoUsecase) Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error) {
	return tu.todoRepository.Update(ctx, input)
}

func (tu todoUsecase) Delete(ctx context.Context, id string) (*model.Todo, error) {
	return tu.todoRepository.Delete(ctx, id)
}
