package repository

import (
	"context"
	"github.com/arangodb/go-driver"
	"log"

	"back/graph/model"
)

type UserRepository interface {
	Create(ctx context.Context, input model.NewUser) (string, error)
	Login(ctx context.Context, input model.Login) (*model.UserData, error)
	GetPlayer(ctx context.Context, playerId string) (model.UserData, error)
}

type userrepository struct {
	db driver.Database
}

func NewUserRepository(db driver.Database) UserRepository {
	return &userrepository{
		db: db,
	}
}

func (ur userrepository) Create(ctx context.Context, input model.NewUser) (string, error) {

	return "", nil
}

func (ur userrepository) GetPlayer(ctx context.Context, playerId string) (model.UserData, error) {
	var player = model.UserData{}

	query := "FOR d IN player FILTER d._id == @playerId RETURN d"
	bindVars := map[string]interface{}{
		"playerId": playerId,
	}

	cursor, err := ur.db.Query(ctx, query, bindVars)
	if err != nil {
		log.Fatal("Error login Query to db")
	}
	defer cursor.Close()

	_, err = cursor.ReadDocument(ctx, &player)

	return player, nil
}

func (ur userrepository) Login(ctx context.Context, input model.Login) (*model.UserData, error) {

	var retuser = &model.UserData{}
	query := "FOR d IN player FILTER d.email == @email RETURN d"
	bindVars := map[string]interface{}{
		"email": input.Username,
	}

	cursor, err := ur.db.Query(ctx, query, bindVars)
	if err != nil {
		log.Fatal("Error login Query to db")
	}
	defer cursor.Close()

	_, err = cursor.ReadDocument(ctx, &retuser)
	return retuser, nil
}
