package controller

import (
	"context"

	"back/graph/model"
	"back/pkg/usecase"
)

type UserController interface {
	Create(ctx context.Context, input model.NewUser) (string, error)
	Login(ctx context.Context, input model.Login) (string, error)
	LoginUser(ctx context.Context, input model.Login) (*model.LoginData, error)
	RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error)
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{
		userUsecase: uu,
	}
}

func (uc userController) Create(ctx context.Context, input model.NewUser) (string, error) {
	return uc.userUsecase.Create(ctx, input)
}

func (uc userController) Login(ctx context.Context, input model.Login) (string, error) {
	return uc.userUsecase.Login(ctx, input)
}

func (uc userController) LoginUser(ctx context.Context, input model.Login) (*model.LoginData, error) {
	return uc.userUsecase.LoginUser(ctx, input)
}

func (uc userController) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	return uc.userUsecase.RefreshToken(ctx, input)
}
