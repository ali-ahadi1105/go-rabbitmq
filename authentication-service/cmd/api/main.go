package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
)

const webPort = "8080"

type Application struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting authentication service...")

	db := connectToDb()
	if db == nil {
		log.Fatal("could not connect to db")
	}

	app := Application{
		DB:     db,
		Models: data.New(db),
	}

	log.Printf("authentication service is running on port: %s\n", webPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}

func connectToDb() *sql.DB {
	dsn := os.Getenv("DSN")

	for range 10 {
		db, err := sql.Open("pgx", dsn)

		if err == nil && db.Ping() == nil {
			log.Fatal("could not connet to database")
			return db
		}

		log.Println("wating for database")
		time.Sleep(time.Second * 2)
	}
	return nil
}
