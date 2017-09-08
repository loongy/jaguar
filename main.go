package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq" // initiliase the postgres driver
	"github.com/loongy/jaguar/actions"
	"github.com/loongy/jaguar/api"
	"github.com/loongy/jaguar/storage/db"
)

func main() {

	// Listening port
	portDefault := os.Getenv("PORT")
	if portDefault == "" {
		portDefault = "5000"
	}
	port := flag.String("port", portDefault, "The port on which the server will listen. Defaults to '5000'.")

	// Database URL
	dbURLDefault := os.Getenv("DATABASE_URL")
	if dbURLDefault == "" {
		dbURLDefault = "host=localhost port=5432 ssl=disabled"
	}
	dbURL := flag.String("db", dbURLDefault, "The database connection URL. Defaults to 'host=localhost port=5432 ssl=disabled'.")

	// Parse flags
	flag.Parse()

	db, err := db.NewDB("postgres", *dbURL)
	if err != nil {
		log.Fatal(err)
	}
	ctx := actions.NewContext(db, nil)
	server, err := api.NewServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port ", *port, "...")
	if err := server.ListenAndServe(":" + *port); err != nil {
		log.Fatal(err)
	}
}
