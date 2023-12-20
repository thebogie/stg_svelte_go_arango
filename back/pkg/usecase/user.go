package usecase

import (
	"back/auth"
	"back/graph/model"
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
	CheckLogin(ctx context.Context, player string) (string, error)
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

func (uu userUsecase) CheckLogin(ctx context.Context, player string) (string, error) {
	status := "WRONGUSER"
	if ctx.Value("authuser").(string) == player {
		status = "OK"
	}

	return status, nil
}

func (uu userUsecase) LoginUser(ctx context.Context, input model.Login) (*model.LoginData, error) {

	founduser, err := uu.userRepository.Login(ctx, input)

	loggedindata := &model.LoginData{
		Token:    "",
		Userdata: founduser,
	}

	jwtheader := ctx.Value("jwtheader").(*auth.JwtHeader)
	if jwtheader.Authenticate(input.Password, founduser.Password) {

		// legit create jwt token

		tokenString, _ := jwtheader.GenerateToken(founduser.Email)
		loggedindata.Token = tokenString

	}

	return loggedindata, err

}

func (uu userUsecase) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	return "", nil
}
