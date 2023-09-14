package db

import (
	"context"
	godriver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"os"
	"time"
)

func InitDB() godriver.Database {
	var err error

	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI == "" {
		databaseURI = "http://localhost:50001"
	}

	var conn godriver.Connection

	log.Printf("DB CONNECTION ENVTORUN SETTING:%v:::DB URI:%v", os.Getenv("ENVTORUN"), databaseURI)

	var uptime = 10
	var client godriver.Client

	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{databaseURI},
	})

	if err != nil {
		log.Fatal("Creation of connection string to failed", err.Error())
	}

	client, err = godriver.NewClient(godriver.ClientConfig{
		Connection:     conn,
		Authentication: godriver.BasicAuthentication("root", "letmein"),
		//Authentication: godriver.BasicAuthentication("root", "wnbGnPpCXHwbP"),
	})
	if err != nil {
		log.Fatal("Creation of NewClient failed", err.Error())
	}

	client.Connection().SetAuthentication(godriver.BasicAuthentication("root", "letmein"))

	for i := 0; i < uptime; i++ {
		log.Printf("Attempting arangodb connection to smacktalk db")
		time.Sleep(5 * time.Second)

		ctx := context.Background()
		huh, err := client.Databases(ctx)
		log.Printf("huh", huh)
		what, err := client.DatabaseExists(ctx, "smacktalk")
		if err != nil {
			log.Printf("Try again. Smacktalk DB isnt found:   ", err.Error())

			if i == uptime {
				log.Fatal("smacktalk DB isnt found after x times", err.Error())
			}
		} else {
			log.Printf("Found smacktalk:", what)
			//worked
			i = 10
		}

	}

	ctx := context.Background()

	huh, err := client.Databases(ctx)
	if err != nil {
		log.Printf("ERROR:", err.Error())
	} else {

		log.Printf("huh:", huh)
	}

	users, err := client.Users(ctx)
	if err != nil {
		log.Printf("ERROR:", err.Error())
	} else {

		log.Printf("users:", users)
	}
	dbca, err := client.Databases(ctx)
	if err != nil {
		log.Printf("DBS ERROR:", err.Error())
	} else {
		log.Printf("DBS:", dbca)
	}

	db, err := client.Database(ctx, "smacktalk")

	if err != nil {
		defer log.Printf("Connection to smacktalk Failed")
		log.Fatal("Smacktalk database not reachable", err.Error())
	} else {
		log.Printf("Connection to Database Successfully")
	}
	DbExceptionHandle(err)

	return db
}

func DbExceptionHandle(err error) {
	if err != nil {
		panic(err)
	}
}
