package repository

import (
	"back/graph/model"
	"back/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"log"
)

type ContestRepository interface {
	List(ctx context.Context) ([]*model.Contest, error)
	GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error)
	CreateContest(ctx context.Context, newContest model.InputContest) (string, error)
	AddVenue(ctx context.Context, newContest *model.InputVenue) (string, error)
	FindVenue(ctx context.Context, newContest *model.InputVenue) (string, error)
}

type contestrepository struct {
	db driver.Database
}

func NewContestRepository(db driver.Database) ContestRepository {
	return &contestrepository{
		db: db,
	}
}

func getNewArangoDocId(queryFunc func(context.Context, string, map[string]interface{}) (driver.Cursor, error), ctx context.Context, queryStr string) string {

	utils.PrintFunctionName()

	cursor, err := queryFunc(ctx, queryStr, nil)
	if err != nil {
		log.Fatalf("Error Query: %s", err)
	}

	var transfer map[string]interface{}
	defer cursor.Close()
	for {
		if !cursor.HasMore() {
			break
		}
		// Decode the result into your struct
		if _, err := cursor.ReadDocument(context.Background(), &transfer); err != nil {
			log.Fatal("Error decoding result:", err)
		}
	}
	return transfer["_id"].(string)
}

// ADD EDGE
func (cr *contestrepository) AddPlayedAt(ctx context.Context, playedAt ArangoDBPlayedAt) (string, error) {
	utils.PrintFunctionName()

	transferJson, err := json.Marshal(playedAt)
	if err != nil {
		fmt.Println("Error marshalling newPlayedAt:", err)
		// Handle error appropriately
	}
	query :=
		`INSERT ` + string(transferJson) + ` INTO played_at RETURN NEW`
	playedAtId := getNewArangoDocId(cr.db.Query, ctx, query)

	return playedAtId, nil
}

func (cr *contestrepository) AddPlayedWith(ctx context.Context, playedWith ArangoDBPlayedWith) (string, error) {
	utils.PrintFunctionName()

	transferJson, err := json.Marshal(playedWith)
	if err != nil {
		fmt.Println("Error marshalling playedWith:", err)
		// Handle error appropriately
	}
	query :=
		`INSERT ` + string(transferJson) + ` INTO played_with RETURN NEW`
	playedWithId := getNewArangoDocId(cr.db.Query, ctx, query)

	return playedWithId, nil
}

func (cr *contestrepository) AddResultedIn(ctx context.Context, resultedIn ArangoDBResultedIn) (string, error) {
	utils.PrintFunctionName()

	transferJson, err := json.Marshal(resultedIn)
	if err != nil {
		fmt.Println("Error marshalling newPlayedAt:", err)
		// Handle error appropriately
	}
	query :=
		`INSERT ` + string(transferJson) + ` INTO resulted_in RETURN NEW`
	resultedInID := getNewArangoDocId(cr.db.Query, ctx, query)

	return resultedInID, nil
}

//END ADD EDGE

func (cr *contestrepository) AddPlayer(ctx context.Context, newPlayer *model.InputUserData) (string, error) {
	utils.PrintFunctionName()

	//assume that the venue coming through exists and find the VenueId IF VenueId is not already there
	playerId, err := cr.FindPlayer(ctx, newPlayer)
	if err != nil {
		//options?
	}

	//TODO: set basic password that has to be reset
	if playerId == "" {
		var arangoPlayer = ArangoDBPlayer{Email: newPlayer.Email, Firstname: newPlayer.Firstname, Password: ""}
		transferJson, err := json.Marshal(arangoPlayer)
		if err != nil {
			fmt.Println("Error marshalling transferJson:", err)
			// Handle error appropriately
		}
		query :=
			`INSERT ` + string(transferJson) + ` INTO player RETURN NEW`
		playerId = getNewArangoDocId(cr.db.Query, ctx, query)
	}

	return playerId, nil
}

func (cr *contestrepository) AddVenue(ctx context.Context, newVenue *model.InputVenue) (string, error) {
	utils.PrintFunctionName()

	//assume that the venue coming through exists and find the VenueId IF VenueId is not already there
	venueId, err := cr.FindVenue(ctx, newVenue)
	if err != nil {
		//options?
	}

	if venueId == "" {
		var arangoVenue = ArangoDBVenue{Address: newVenue.Address, Lat: newVenue.Lat, Lng: newVenue.Lat}
		transferJson, err := json.Marshal(arangoVenue)
		if err != nil {
			fmt.Println("Error marshalling contestJson:", err)
			// Handle error appropriately
		}
		query :=
			`INSERT ` + string(transferJson) + ` INTO venue RETURN NEW`
		venueId = getNewArangoDocId(cr.db.Query, ctx, query)
	}

	return venueId, nil
}

func (cr *contestrepository) AddOutcomes(ctx context.Context, contestID string, newOutcomes []*model.InputOutcome) {
	utils.PrintFunctionName()

	for _, outcome := range newOutcomes {
		playerId, err := cr.FindPlayer(ctx, outcome.Player)
		if err != nil {
			//options?
		}

		var resultedInObj = ArangoDBResultedIn{Label: "RESULTED_IN", Place: outcome.Place, From: contestID, To: playerId, Result: outcome.Result}
		_, err = cr.AddResultedIn(ctx, resultedInObj)
		if err != nil {
			//options?
		}

	}

	return
}

func (cr *contestrepository) AddGames(ctx context.Context, contestID string, newGames []*model.InputGame) {
	utils.PrintFunctionName()

	for _, game := range newGames {
		gameID, err := cr.FindGame(ctx, game)
		if err != nil {
			//options?
		}

		var playedWithObj = ArangoDBPlayedWith{From: contestID, To: gameID, Label: "PLAYED_WITH"}
		_, err = cr.AddPlayedWith(ctx, playedWithObj)
		if err != nil {
			//options?
		}

	}

	return
}

func (cr *contestrepository) FindVenue(ctx context.Context, findvenue *model.InputVenue) (string, error) {
	utils.PrintFunctionName()
	venueKey := findvenue.Key

	query := `
FOR doc IN venue
FILTER (UPPER(doc.address) == UPPER("` + findvenue.Address +
		`")) OR ((doc.lat == "` + findvenue.Lng +
		`") AND (doc.lng == "` + findvenue.Lat + `"))
RETURN doc
`

	if findvenue.Key == "" {
		cursor, err := cr.db.Query(ctx, query, nil)
		if err != nil {
			log.Fatalf("Error Query: %s", err)
		}

		defer cursor.Close()
		for {
			if !cursor.HasMore() {
				break
			}

			var transfer map[string]interface{}

			// Decode the result into your struct
			if _, err := cursor.ReadDocument(context.Background(), &transfer); err != nil {
				log.Fatal("Error decoding result:", err)
			}

			venueKey = transfer["_id"].(string)
		}
	}

	// upsert venue
	return venueKey, nil
}

func (cr *contestrepository) FindPlayer(ctx context.Context, player *model.InputUserData) (string, error) {
	utils.PrintFunctionName()
	playerId := player.ID

	query := `
FOR doc IN player
FILTER (UPPER(doc.email) == UPPER("` + player.Email +
		`"))
RETURN doc
`

	if playerId == "" {
		playerId = getNewArangoDocId(cr.db.Query, ctx, query)
	}

	// upsert venue
	return playerId, nil
}

func (cr *contestrepository) FindGame(ctx context.Context, game *model.InputGame) (string, error) {
	utils.PrintFunctionName()
	gameID := game.ID

	query := `
FOR doc IN game
FILTER (UPPER(doc.name) == UPPER("` + game.Name +
		`"))
RETURN doc
`

	if gameID == "" {
		gameID = getNewArangoDocId(cr.db.Query, ctx, query)
	}

	// upsert venue
	return gameID, nil
}

func (cr *contestrepository) CreateContest(ctx context.Context, newContest model.InputContest) (string, error) {
	utils.PrintFunctionName()

	contestName := utils.GetContestName()

	var transferContest = ArangoDBContest{Name: contestName, Start: newContest.Start, Stop: newContest.Stop, Startoffset: newContest.Startoffset, Stopoffset: newContest.Stopoffset}
	var newContestID = ""

	transferJson, err := json.Marshal(transferContest)
	if err != nil {
		fmt.Println("Error marshalling newContest:", err)
		// Handle error appropriately
	}
	query :=
		`INSERT ` + string(transferJson) + ` INTO contest RETURN NEW`
	newContestID = getNewArangoDocId(cr.db.Query, ctx, query)

	cr.AddOutcomes(ctx, newContestID, newContest.Outcomes)

	venueId, err := cr.AddVenue(ctx, newContest.Venue)
	if err != nil {
		//options?
	}

	var playedAtObj = ArangoDBPlayedAt{Label: "PLAYED_AT", From: newContestID, To: venueId}
	_, err = cr.AddPlayedAt(ctx, playedAtObj)
	if err != nil {
		//options?
	}

	cr.AddGames(ctx, newContestID, newContest.Games)

	return newContestID, nil
}

func (cr *contestrepository) List(ctx context.Context) ([]*model.Contest, error) {
	//var retuser = &model.UserData{}

	return []*model.Contest{}, nil
}

func (cr *contestrepository) GetContestsPlayerTotalResults(ctx context.Context, player string) ([]*model.Contest, error) {
	utils.PrintFunctionName()
	/* query := "" +
	"FOR doc IN resulted_in " + " " +
	"FILTER doc._to == '" + player + "'" + " " +
	"RETURN { contest: doc._from, results: {player: doc._to, place: doc.place, result: doc.result }}"
	*/
	//query := "FOR contest IN contest\n   LET player_resulted_in = (\n     FOR resulted_in IN resulted_in\n       FILTER resulted_in._from == contest._id AND resulted_in._to == \"player/202312291853567769250600\"\n       RETURN resulted_in\n   )\n\n\n   FILTER LENGTH(player_resulted_in) > 0\n   RETURN {\n      played_with: (\n       FOR played_with IN played_with\n         FILTER played_with._from == contest._id\n           FOR to_doc IN game\n   FILTER to_doc._id == played_with._to\n   RETURN to_doc.name\n     ),\n\n\n\n          played_at: (\n       FOR played_at IN played_at\n         FILTER played_at._from == contest._id\n                   FOR to_doc IN venue\n   FILTER to_doc._id == played_at._to\n   RETURN to_doc.address\n     ),\n     contest: contest,\n     resulted_in: (\n       FOR resulted_in IN resulted_in\n         FILTER resulted_in._from == contest._id\n         RETURN resulted_in\n     )\n   }"

	query :=
		`
	FOR
	contest
	IN
	contest
	LET
	player_resulted_in = (
		FOR
	resulted_in
	IN
	resulted_in
	FILTER
	resulted_in._from == contest._id
	AND
	resulted_in._to == "` + player + `"
	RETURN
	resulted_in
	)
	FILTER
	LENGTH(player_resulted_in) > 0
	RETURN{
		played_with: (
			FOR played_with IN played_with
		FILTER played_with._from == contest._id
		FOR to_doc IN game
		FILTER to_doc._id == played_with._to
		RETURN to_doc
	),
		played_at: (
		FOR played_at IN played_at
		FILTER played_at._from == contest._id
		FOR to_doc IN venue
		FILTER to_doc._id == played_at._to
		RETURN to_doc
	),
		contest: contest,
		resulted_in: (
		FOR resulted_in IN resulted_in
		FILTER resulted_in._from == contest._id
		RETURN resulted_in
	)
	}
	`

	cursor, err := cr.db.Query(ctx, query, nil)
	if err != nil {
		log.Fatalf("Error Query: %s", err)
	}

	defer cursor.Close()
	var results = []*model.Contest{}
	//var results_allcontests []model.Contest

	for {
		if !cursor.HasMore() {
			break
		}

		var played_at = &model.Venue{}
		var played_with = []*model.Game{}
		var addContest = &model.Contest{}
		var resulted_in = []*model.Outcome{}
		var transfer map[string]interface{}

		// Decode the result into your struct
		if _, err := cursor.ReadDocument(context.Background(), &transfer); err != nil {
			log.Fatal("Error decoding result:", err)
		}

		//played_at = Venue
		value, ok := transfer["played_at"]
		if !ok {
			// Handle the case where the "played_at" key is not found
			// ...
		} else {
			playedAt, ok := value.([]interface{})
			if !ok {
				log.Printf("Handle the case where the value is not a slice")
				// ...
			} else if len(playedAt) == 0 {
				log.Printf("Handle the case where the slice is empty ")
				// ...
			} else {
				// Now you can safely iterate over the first element:
				jsonBytes, err := json.Marshal(playedAt[0])
				if err != nil {
					log.Printf("Error with marshalling playedAt")
				}
				err = json.Unmarshal(jsonBytes, &played_at)
				if err != nil {
					log.Printf("Error with unmarshalling playedAt", err)
				}

			}
		}

		//played_with = slice of Games
		value, ok = transfer["played_with"]
		if !ok {
			// Handle the case where the "played_at" key is not found
			// ...
		} else {
			playedWith, ok := value.([]interface{})
			if !ok {
				log.Printf("Handle the case where the value is not a slice")
				// ...
			} else if len(playedWith) == 0 {
				log.Printf("Handle the case where the slice is empty ")
				// ...
			} else {
				// Now you can safely iterate over the elements:
				for i, _ := range playedWith {
					jsonBytes, err := json.Marshal(playedWith[i])
					if err != nil {
						log.Printf("Error with marshalling playedAt")
					}
					game_element := model.Game{}
					err = json.Unmarshal(jsonBytes, &game_element)
					if err != nil {
						log.Printf("Error with unmarshalling playedAt %s", err)
					}
					played_with = append(played_with, &game_element)
				}

			}
		}

		//resulted_in = slice of Outcome
		value, ok = transfer["resulted_in"]
		if !ok {
			// Handle the case where the "played_at" key is not found
			// ...
		} else {
			resultedIn, ok := value.([]interface{})
			if !ok {
				log.Printf("Handle the case where the value is not a slice")
				// ...
			} else if len(resultedIn) == 0 {
				log.Printf("Handle the case where the slice is empty ")
				// ...
			} else {
				// Now you can safely iterate over the elements:
				for i, item := range resultedIn {

					jsonBytes, err := json.Marshal(resultedIn[i])
					if err != nil {
						log.Printf("Error with marshalling playedAt")
					}
					outcome_element := model.Outcome{}
					err = json.Unmarshal(jsonBytes, &outcome_element)
					if err != nil {
						log.Printf("Error with unmarshalling playedAt %s", err)
					}

					//convert _to to Player
					if mapItem, ok := item.(map[string]interface{}); ok {
						// It's a map, extract the key-value pair

						repo, _ := NewUserRepository(cr.db).GetPlayer(ctx, mapItem["_to"].(string))

						outcome_element.Player = &repo
						// clear password...
						outcome_element.Player.Password = ""

					}

					resulted_in = append(resulted_in, &outcome_element)
				}

			}

		}

		//contest = Contest
		contestObj, ok := transfer["contest"]
		jsonBytes, err := json.Marshal(contestObj)
		if err != nil {
			log.Printf("Error with marshalling playedAt")
		}
		err = json.Unmarshal(jsonBytes, &addContest)
		if err != nil {
			log.Printf("Error with unmarshalling playedAt", err)
		}

		addContest.Games = played_with
		addContest.Venue = played_at
		addContest.Outcomes = resulted_in
		//TODO: fix contest id, rev and Player from _to in resulted_in
		results = append(results, addContest)

	}

	// Process the result
	log.Printf("Results: %+v\n", results)
	//results = append(results, &result)

	return results, nil
}
