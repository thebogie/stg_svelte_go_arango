package repository

import (
	"back/graph/model"
	"context"
	"github.com/arangodb/go-driver"
	"log"
)

type GameRepository interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Game, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
}

type gamerepository struct {
	db driver.Database
}

func NewGameRepository(db driver.Database) GameRepository {
	return &gamerepository{
		db: db,
	}
}

func (gr *gamerepository) List(ctx context.Context) ([]*model.Game, error) {
	query := "FOR doc IN game RETURN doc"
	//TODO: pass context from top level?

	cursor, err := gr.db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal("Error login Query to db")
	}
	defer cursor.Close()

	//_, err = cursor.ReadDocument(ctx, &retuser)
	var results []*model.Game // Replace with your struct type

	// Iterate over the cursor and append results to the array
	for {
		var doc model.Game // Replace with your struct type
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break // No more documents in the cursor
		} else if err != nil {
			log.Fatal(err)
		}
		results = append(results, &doc)
	}

	return results, nil
}

//func (tr *gamerepository) ListPlayedByPlayerid(ctx context.Context) ([]*model.Game, error) {
//query := "For c in contest LET involvedfull = (    FOR v, e IN 1..3 INBOUND  @playerid resulted_in    RETURN DISTINCT v    )    for findgame in  involvedfull  FOR v, e IN 1..3 OUTBOUND findgame played_with  RETURN DISTINCT v"
//bindVars := map[string]interface{}{
//	"playerid": playerid,
//}

//return []*model.Game{}, nil
//}
