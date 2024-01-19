package repository

import (
	"back/graph/model"
	"context"
	"github.com/arangodb/go-driver"
	"log"
)

type GameRepository interface {
	List(ctx context.Context) ([]*model.Game, error)
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
