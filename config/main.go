package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"golang.org/x/crypto/bcrypt"
)

type Outcome struct {
	Key      string `json:"_key"`
	Playerid string `json:"playerid"`
	Place    string `json:"place"`
	Result   string `json:"result"`
}

type Game struct {
	Key  string `json:"_key"`
	Name string `json:"name"`
}

type Venue struct {
	Key     string `json:"_key"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

type Player struct {
	Key  string `json:"_key"`
	Name string `json:"playerid"`
}

type Played_At struct {
	Key   string `json:"_key"`
	To    string `json:"_to"`
	From  string `json:"_from"`
	Label string `json:"_label"`
}

type Played_With struct {
	Key   string `json:"_key"`
	To    string `json:"_to"`
	From  string `json:"_from"`
	Label string `json:"_label"`
}

type Resulted_In struct {
	Key    string `json:"_key"`
	To     string `json:"_to"`
	From   string `json:"_from"`
	Label  string `json:"_label"`
	Place  string `json:"place"`
	Result string `json:"result"`
}

type Contest struct {
	Key         string `json:"_key"`
	Start       string `json:"start"`
	Startoffset string `json:"startoffset"`
	Stop        string `json:"stop"`
	Stopoffset  string `json:"stopoffset"`
	Venue       Venue
	Outcome     []Outcome
	Games       []string
}

type EntityContest struct {
	Key         string `json:"_key"`
	Start       string `json:"start"`
	Startoffset string `json:"startoffset"`
	Stop        string `json:"stop"`
	Stopoffset  string `json:"stopoffset"`
}

type EntityPlayer struct {
	Key       string `json:"_key"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// collections
type iBase interface {
	findkey()
}

type Venues struct {
	VenueList []Venue
}

func (b *Venues) findkey(find string) string {

	retval := ""
	for _, x := range b.VenueList {

		if x.Address == find {
			retval = x.Key
		}
	}

	return retval

}

type Games struct {
	GameList []Game
}

func (b *Games) findkey(find string) string {

	retval := ""
	for _, x := range b.GameList {

		if x.Name == find {
			retval = x.Key
		}
	}

	return retval

}

type Players struct {
	PlayerList []Player
}

func (b *Players) findkey(find string) string {

	retval := ""
	for _, x := range b.PlayerList {

		if x.Name == find {
			retval = x.Key
		}
	}

	return retval

}

func main() {

	var err error
	var client driver.Client
	var conn driver.Connection

	var reset bool
	flag.BoolVar(&reset, "r", false, "reset database")
	flag.Parse()

	//reset = true
	log.Printf("Reset Db? %+v", reset)

	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:50001"},
		//Endpoints: []string{"http://192.168.86.222:8529"},
		//Endpoints: []string{"https://5a812333269f.arangodb.cloud:8529/"},
	})

	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}
	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "letmein"),
		//Authentication: driver.BasicAuthentication("root", "wnbGnPpCXHwbP"),
	})

	if err != nil {
		log.Fatalf("Failed to connect to db: %+v", err)
	}

	var db driver.Database
	var db_exists bool
	//var coll_exists bool

	db_exists, err = client.DatabaseExists(nil, "smacktalk")

	if db_exists && reset {

		db, err = client.Database(nil, "smacktalk")

		err = db.Remove(nil)

		if err != nil {
			log.Fatalf("Failed to drop database: %+v", err)
		}
	}

	db_exists, err = client.DatabaseExists(nil, "smacktalk")

	if db_exists {
		log.Println("That db exists already")

		db, err = client.Database(nil, "smacktalk")

		if err != nil {
			log.Fatalf("Failed to open existing database: %v", err)
		}

	} else {
		db, err = client.CreateDatabase(nil, "smacktalk", nil)

		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
	}

	//read in json
	log.Println("READINJSON")

	var contests []Contest
	var venues Venues
	var games Games
	var players Players

	var finalplayers []EntityPlayer
	var finalcontests []EntityContest

	var playedats []Played_At
	var playedwiths []Played_With
	var resultedins []Resulted_In

	filecontent, err := os.Open("stg_records.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer filecontent.Close()

	byteResult, _ := ioutil.ReadAll(filecontent)

	json.Unmarshal(byteResult, &contests)

	for index := range contests {

		//find unique venues
		var exists bool
		exists = false
		for _, x := range venues.VenueList {
			if x.Address == contests[index].Venue.Address {
				exists = true
				break
			}
		}

		if !exists {

			contests[index].Venue.Key = gettimestamp()
			venues.VenueList = append(venues.VenueList, contests[index].Venue)
		}

		//find unique games

		var gamename string
		for _, y := range contests[index].Games {
			exists = false
			for _, x := range games.GameList {

				gamename = y
				if x.Name == y {
					exists = true
					break
				}
			}
			if !exists {
				var addgame Game

				addgame.Key = gettimestamp()
				addgame.Name = gamename
				games.GameList = append(games.GameList, addgame)
			}
		}

		//find unique players
		var playername string
		for _, y := range contests[index].Outcome {
			exists = false
			for _, x := range players.PlayerList {

				playername = y.Playerid
				if x.Name == y.Playerid {
					exists = true
					break
				}
			}

			if !exists {
				var addplayer Player

				addplayer.Key = gettimestamp()
				addplayer.Name = playername
				players.PlayerList = append(players.PlayerList, addplayer)
			}
		}

	}

	for _, contest := range contests {

		var finalcontest EntityContest

		finalcontest.Key = gettimestamp()
		finalcontest.Start = contest.Start
		finalcontest.Startoffset = contest.Startoffset
		finalcontest.Stop = contest.Stop
		finalcontest.Stopoffset = contest.Stopoffset

		var playedat Played_At
		playedat.Key = gettimestamp()
		playedat.From = "contest/" + finalcontest.Key
		playedat.To = "venue/" + venues.findkey(contest.Venue.Address)
		playedat.Label = "PLAYED_AT"
		playedats = append(playedats, playedat)

		for _, g := range contest.Games {
			var playedwith Played_With
			playedwith.Key = gettimestamp()
			playedwith.From = "contest/" + finalcontest.Key
			playedwith.To = "game/" + games.findkey(g)
			playedwith.Label = "PLAYED_WITH"
			playedwiths = append(playedwiths, playedwith)

		}

		for _, r := range contest.Outcome {
			var resultedin Resulted_In
			resultedin.Key = gettimestamp()
			resultedin.From = "contest/" + finalcontest.Key
			resultedin.To = "player/" + players.findkey(r.Playerid)
			resultedin.Label = "RESULTED_IN"
			resultedin.Place = r.Place
			resultedin.Result = r.Result
			resultedins = append(resultedins, resultedin)

		}
		finalcontests = append(finalcontests, finalcontest)

	}

	//Convert players into entity
	for _, player := range players.PlayerList {

		var finalplayer EntityPlayer

		finalplayer.Key = player.Key
		finalplayer.Firstname = player.Name
		finalplayer.Email = player.Name + "@gmail.com"
		finalplayer.Password = HashPassword("letmein")

		if player.Name != "" {
			finalplayers = append(finalplayers, finalplayer)
		}

	}

	//log.Println("GAMES: %v", games)
	gamesJSON, _ := json.Marshal(games.GameList)
	log.Println("Creating games")
	err = ioutil.WriteFile("stg_games.json", gamesJSON, 0777)

	//log.Println("VENUES: %v", venues.VenueList)
	venuesJSON, _ := json.Marshal(venues.VenueList)
	log.Println("Creating venues")
	err = ioutil.WriteFile("stg_venues.json", venuesJSON, 0777)

	//log.Println("PLAYERS: %v", players)
	playersJSON, _ := json.Marshal(finalplayers)
	log.Println("Creating players")
	err = ioutil.WriteFile("stg_players.json", playersJSON, 0777)

	//log.Println("FINAL: %v", finalcontests)
	contestJSON, _ := json.Marshal(finalcontests)
	log.Println("Creating contests")
	err = ioutil.WriteFile("stg_contests.json", contestJSON, 0777)

	//log.Println("PLAYEDAT: %v", playedats)
	playedatJSON, _ := json.Marshal(playedats)
	log.Println("Creating played at")
	err = ioutil.WriteFile("stg_playedat.json", playedatJSON, 0777)

	//log.Println("PLAYEWITH: %v", playwiths)
	playedwithJSON, _ := json.Marshal(playedwiths)
	log.Println("Creating played with")
	err = ioutil.WriteFile("stg_playedwith.json", playedwithJSON, 0777)

	//log.Println("RESULTEDIN: %v", resultedins)
	resultedinJSON, _ := json.Marshal(resultedins)
	log.Println("Creating resulted in")
	err = ioutil.WriteFile("stg_resultedin.json", resultedinJSON, 0777)

	//load it all into arango
	loadtoarango(db, "player", finalplayers, 2)
	loadtoarango(db, "game", games.GameList, 2)
	loadtoarango(db, "venue", venues.VenueList, 2)
	loadtoarango(db, "contest", finalcontests, 2)

	loadtoarango(db, "played_at", playedats, 3)
	loadtoarango(db, "played_with", playedwiths, 3)
	loadtoarango(db, "resulted_in", resultedins, 3)

}

func gettimestamp() string {

	time.Sleep(time.Duration(rand.Intn(100-10)+10) * time.Nanosecond)
	retval := strings.Replace((time.Now().Format(time.RFC3339Nano)), ":", "", -1)
	retval = strings.Replace(retval, "-", "", -1)
	retval = strings.Replace(retval, ".", "", -1)
	retval = strings.Replace(retval, "T", "", -1)

	return retval
}

func loadtoarango(db driver.Database, colname string, list interface{}, typeofdoc int) {

	var col driver.Collection
	ctx := context.Background()
	options := &driver.CreateCollectionOptions{}
	if typeofdoc == 2 {
		options.Type = driver.CollectionTypeDocument
	}
	if typeofdoc == 3 {
		options.Type = driver.CollectionTypeEdge
	}

	col, err := db.CreateCollection(ctx, colname, options)
	if err != nil {

	}
	_, errs, err := col.CreateDocuments(nil, list)
	if err != nil {
		log.Fatalf("Failed to create documents: %v", err)
	} else if err := errs.FirstNonNil(); err != nil {
		log.Fatalf("Failed to create documents: first errornote %v", err)
	}

}

func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to create documents: %v", err.Error())
	}
	return string(result)
}
