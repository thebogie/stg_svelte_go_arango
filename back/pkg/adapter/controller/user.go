package controller

import (
	"back/graph/model"
	"back/pkg/usecase"
	"context"
)

type UserController interface {
	Create(ctx context.Context, input model.NewUser) (string, error)
	Login(ctx context.Context, input model.Login) (string, error)
	LoginUser(ctx context.Context, input model.Login) (*model.LoginData, error)
	CheckLogin(ctx context.Context, player string) (string, error)
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
func (uc userController) CheckLogin(ctx context.Context, player string) (string, error) {

	return uc.userUsecase.CheckLogin(ctx, player)
}
func (uc userController) LoginUser(ctx context.Context, input model.Login) (*model.LoginData, error) {
	data, err := uc.userUsecase.LoginUser(ctx, input)

	if err != nil || data.Token == "" {
		//ctxw := ctx.Value("ResponseWriter").(http.ResponseWriter)
		data.Userdata.Password = ""
		//TODO: fix how we throw error
		//http.Error(ctx, "Forbidden", http.StatusForbidden)
	}

	return data, err
}

func (uc userController) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {

	return uc.userUsecase.RefreshToken(ctx, input)
}
