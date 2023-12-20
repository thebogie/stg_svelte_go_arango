package repository

import (
	"back/graph/model"
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
	"log"
	"strconv"
)

type ContestRepository interface {
	//Get(ctx context.Context, id string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Contest, error)
	GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error)
	//Create(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	//Update(ctx context.Context, input model.CompletedstatusInput) (*model.Todo, error)
	//Delete(ctx context.Context, id string) (*model.Todo, error)
}

type contestrepository struct {
	db driver.Database
}

func NewContestRepository(db driver.Database) ContestRepository {
	return &contestrepository{
		db: db,
	}
}

func (cr *contestrepository) List(ctx context.Context) ([]*model.Contest, error) {
	//var retuser = &model.UserData{}

	return []*model.Contest{}, nil
}

func (cr *contestrepository) GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error) {
	query := "" +
		"FOR doc IN resulted_in " + " " +
		"FILTER doc._to == '" + player + "'" + " " +
		"RETURN { contest: doc._from, results: {player: doc._to, place: doc.place, result: doc.result }}"
	cursor, err := cr.db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal("Error login Query to db")
	}

	defer cursor.Close()
	var results []*model.Contest
	for {
		if !cursor.HasMore() {
			break
		}
		var result model.Contest
		var outcome model.Outcome
		var transfer map[string]interface{}
		// Decode the result into your struct
		if _, err := cursor.ReadDocument(context.Background(), &transfer); err != nil {
			log.Fatal("Error decoding result:", err)
		}

		result.ID = transfer["contest"].(string)
		outcomevalue, _ := transfer["results"].(map[string]interface{})

		var i, _ = strconv.Atoi(outcomevalue["place"].(string))
		outcome.Player = outcomevalue["player"].(string)
		outcome.Place = i
		outcome.Result = outcomevalue["result"].(string)
		result.Outcomes = append(result.Outcomes, &outcome)

		// Process the result
		fmt.Printf("Result: %+v\n", result)
		results = append(results, &result)
	}

	return results, nil
}
