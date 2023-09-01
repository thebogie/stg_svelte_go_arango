package usecase

import (
	"back/auth"
	"back/graph/model"
	"back/helper"
	"back/pkg/adapter/repository"
	"context"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	Create(ctx context.Context, input model.NewUser) (string, error)
	Login(ctx context.Context, input model.Login) (string, error)
	LoginUser(ctx context.Context, input model.Login) (*model.LoginData, error)
	RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error)
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu userUsecase) Create(ctx context.Context, input model.NewUser) (string, error) {
	return uu.userRepository.Create(ctx, input)
}

func (uu userUsecase) Login(ctx context.Context, input model.Login) (string, error) {
	uu.userRepository.Login(ctx, input)
	return "", nil
}

func (uu userUsecase) LoginUser(ctx context.Context, input model.Login) (*model.LoginData, error) {

	founduser, _ := uu.userRepository.Login(ctx, input)

	loggedindata := &model.LoginData{
		Token:    "",
		Userdata: founduser,
	}

	if helper.Authenticate(input.Password, founduser.Password) {
		// legit create jwt cookie
		cookieaccess := auth.CookieAccess{
			HttpReader: ctx.Value("cookiemaker").(*auth.CookieAccess).HttpReader,
			Writer:     ctx.Value("cookiemaker").(*auth.CookieAccess).Writer,
		}

		tokenString, _ := cookieaccess.GenerateToken(founduser.Email)

		//TODO: what if cookie fails?
		cookieaccess.GenerateAuthCookie(tokenString)
		loggedindata.Token = tokenString
		return loggedindata, nil
	} else {

		return nil, nil
	}
}

func (uu userUsecase) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	return "", nil
}
